package deliverymq

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hookdeck/outpost/internal/alert"
	"github.com/hookdeck/outpost/internal/backoff"
	"github.com/hookdeck/outpost/internal/consumer"
	"github.com/hookdeck/outpost/internal/destregistry"
	"github.com/hookdeck/outpost/internal/idempotence"
	"github.com/hookdeck/outpost/internal/logging"
	"github.com/hookdeck/outpost/internal/models"
	"github.com/hookdeck/outpost/internal/mqs"
	"github.com/hookdeck/outpost/internal/redis"
	"github.com/hookdeck/outpost/internal/scheduler"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

func idempotencyKeyFromDeliveryEvent(deliveryEvent models.DeliveryEvent) string {
	return "idempotency:deliverymq:" + deliveryEvent.ID
}

var (
	errDestinationDisabled = errors.New("destination disabled")
)

// Error types to distinguish between different stages of delivery
type PreDeliveryError struct {
	err error
}

func (e *PreDeliveryError) Error() string {
	return fmt.Sprintf("pre-delivery error: %v", e.err)
}

func (e *PreDeliveryError) Unwrap() error {
	return e.err
}

type DeliveryError struct {
	err error
}

func (e *DeliveryError) Error() string {
	return fmt.Sprintf("delivery error: %v", e.err)
}

func (e *DeliveryError) Unwrap() error {
	return e.err
}

type PostDeliveryError struct {
	err error
}

func (e *PostDeliveryError) Error() string {
	return fmt.Sprintf("post-delivery error: %v", e.err)
}

func (e *PostDeliveryError) Unwrap() error {
	return e.err
}

type messageHandler struct {
	eventTracer    DeliveryTracer
	logger         *logging.Logger
	logMQ          LogPublisher
	entityStore    DestinationGetter
	logStore       EventGetter
	retryScheduler RetryScheduler
	retryBackoff   backoff.Backoff
	retryMaxLimit  int
	idempotence    idempotence.Idempotence
	publisher      Publisher
	alertMonitor   AlertMonitor
}

type Publisher interface {
	PublishEvent(ctx context.Context, destination *models.Destination, event *models.Event) (*models.Delivery, error)
}

type LogPublisher interface {
	Publish(ctx context.Context, deliveryEvent models.DeliveryEvent) error
}

type RetryScheduler interface {
	Schedule(ctx context.Context, task string, delay time.Duration, opts ...scheduler.ScheduleOption) error
	Cancel(ctx context.Context, taskID string) error
}

type DestinationGetter interface {
	RetrieveDestination(ctx context.Context, tenantID, destID string) (*models.Destination, error)
}

type EventGetter interface {
	RetrieveEvent(ctx context.Context, tenantID, eventID string) (*models.Event, error)
}

type DeliveryTracer interface {
	Deliver(ctx context.Context, deliveryEvent *models.DeliveryEvent, destination *models.Destination) (context.Context, trace.Span)
}

type AlertMonitor interface {
	HandleAttempt(ctx context.Context, attempt alert.DeliveryAttempt) error
}

func NewMessageHandler(
	logger *logging.Logger,
	redisClient redis.Cmdable,
	logMQ LogPublisher,
	entityStore DestinationGetter,
	logStore EventGetter,
	publisher Publisher,
	eventTracer DeliveryTracer,
	retryScheduler RetryScheduler,
	retryBackoff backoff.Backoff,
	retryMaxLimit int,
	alertMonitor AlertMonitor,
) consumer.MessageHandler {
	return &messageHandler{
		eventTracer:    eventTracer,
		logger:         logger,
		logMQ:          logMQ,
		entityStore:    entityStore,
		logStore:       logStore,
		publisher:      publisher,
		retryScheduler: retryScheduler,
		retryBackoff:   retryBackoff,
		retryMaxLimit:  retryMaxLimit,
		idempotence: idempotence.New(redisClient,
			idempotence.WithTimeout(5*time.Second),
			idempotence.WithSuccessfulTTL(24*time.Hour),
		),
		alertMonitor: alertMonitor,
	}
}

func (h *messageHandler) Handle(ctx context.Context, msg *mqs.Message) error {
	deliveryEvent := models.DeliveryEvent{}

	// Parse message
	if err := deliveryEvent.FromMessage(msg); err != nil {
		return h.handleError(msg, &PreDeliveryError{err: err})
	}

	h.logger.Ctx(ctx).Info("processing delivery event",
		zap.String("delivery_event_id", deliveryEvent.ID),
		zap.String("event_id", deliveryEvent.Event.ID),
		zap.Int("attempt", deliveryEvent.Attempt))

	// Ensure event data
	if err := h.ensureDeliveryEvent(ctx, &deliveryEvent); err != nil {
		return h.handleError(msg, &PreDeliveryError{err: err})
	}

	// Get destination
	destination, err := h.ensurePublishableDestination(ctx, deliveryEvent)
	if err != nil {
		return h.handleError(msg, &PreDeliveryError{err: err})
	}

	// Handle delivery
	err = h.idempotence.Exec(ctx, idempotencyKeyFromDeliveryEvent(deliveryEvent), func(ctx context.Context) error {
		return h.doHandle(ctx, deliveryEvent, destination)
	})
	return h.handleError(msg, err)
}

func (h *messageHandler) handleError(msg *mqs.Message, err error) error {
	shouldNack := h.shouldNackError(err)
	if shouldNack {
		msg.Nack()
	} else {
		msg.Ack()
	}

	// Don't return error for expected cases
	var preErr *PreDeliveryError
	if errors.As(err, &preErr) {
		if errors.Is(preErr.err, models.ErrDestinationDeleted) || errors.Is(preErr.err, errDestinationDisabled) {
			return nil
		}
	}
	return err
}

func (h *messageHandler) doHandle(ctx context.Context, deliveryEvent models.DeliveryEvent, destination *models.Destination) error {
	_, span := h.eventTracer.Deliver(ctx, &deliveryEvent, destination)
	defer span.End()

	delivery, err := h.publisher.PublishEvent(ctx, destination, &deliveryEvent.Event)
	if err != nil {
		// If delivery is nil, it means no delivery was made.
		// This is an unexpected error and considered a pre-delivery error.
		if delivery == nil {
			return &PreDeliveryError{err: err}
		}

		h.logger.Ctx(ctx).Error("failed to publish event",
			zap.Error(err),
			zap.String("delivery_event_id", deliveryEvent.ID),
			zap.String("destination_id", destination.ID))
		deliveryErr := &DeliveryError{err: err}

		if h.shouldScheduleRetry(deliveryEvent, err) {
			if retryErr := h.scheduleRetry(ctx, deliveryEvent); retryErr != nil {
				return h.logDeliveryResult(ctx, &deliveryEvent, destination, delivery, errors.Join(err, retryErr))
			}
		}
		return h.logDeliveryResult(ctx, &deliveryEvent, destination, delivery, deliveryErr)
	}

	// Handle successful delivery
	if deliveryEvent.Manual {
		logger := h.logger.Ctx(ctx)
		if err := h.retryScheduler.Cancel(ctx, deliveryEvent.GetRetryID()); err != nil {
			h.logger.Ctx(ctx).Error("failed to cancel scheduled retry",
				zap.Error(err),
				zap.String("delivery_event_id", deliveryEvent.ID),
				zap.String("retry_id", deliveryEvent.GetRetryID()))
			return h.logDeliveryResult(ctx, &deliveryEvent, destination, delivery, err)
		}
		logger.Audit("scheduled retry canceled",
			zap.String("delivery_event_id", deliveryEvent.ID),
			zap.String("retry_id", deliveryEvent.GetRetryID()))
	}
	return h.logDeliveryResult(ctx, &deliveryEvent, destination, delivery, nil)
}

func (h *messageHandler) logDeliveryResult(ctx context.Context, deliveryEvent *models.DeliveryEvent, destination *models.Destination, delivery *models.Delivery, err error) error {
	logger := h.logger.Ctx(ctx)

	// Set up delivery record
	deliveryEvent.Delivery = delivery

	logger.Audit("event delivered",
		zap.String("delivery_event_id", deliveryEvent.ID),
		zap.String("destination_id", deliveryEvent.DestinationID),
		zap.String("event_id", deliveryEvent.Event.ID),
		zap.String("delivery_status", deliveryEvent.Delivery.Status),
		zap.Int("attempt", deliveryEvent.Attempt),
		zap.Bool("manual", deliveryEvent.Manual),
		zap.String("destination_type", destination.Type))

	// Publish delivery log
	if logErr := h.logMQ.Publish(ctx, *deliveryEvent); logErr != nil {
		logger.Error("failed to publish delivery log",
			zap.Error(logErr),
			zap.String("delivery_event_id", deliveryEvent.ID))
		if err != nil {
			return &PostDeliveryError{err: errors.Join(err, logErr)}
		}
		return &PostDeliveryError{err: logErr}
	}

	// Call alert monitor in goroutine
	go h.handleAlertAttempt(ctx, deliveryEvent, destination, err)

	// If we have a DeliveryError, return it as is
	var delErr *DeliveryError
	if errors.As(err, &delErr) {
		return err
	}

	// If we have a PreDeliveryError, return it as is
	var preErr *PreDeliveryError
	if errors.As(err, &preErr) {
		return err
	}

	// For any other error, wrap it in PostDeliveryError
	if err != nil {
		return &PostDeliveryError{err: err}
	}

	return nil
}

func (h *messageHandler) handleAlertAttempt(ctx context.Context, deliveryEvent *models.DeliveryEvent, destination *models.Destination, err error) {
	attempt := alert.DeliveryAttempt{
		Success:       deliveryEvent.Delivery.Status == models.DeliveryStatusSuccess,
		DeliveryEvent: deliveryEvent,
		Destination: &alert.AlertDestination{
			ID:         destination.ID,
			TenantID:   destination.TenantID,
			Type:       destination.Type,
			Topics:     destination.Topics,
			Config:     destination.Config,
			CreatedAt:  destination.CreatedAt,
			DisabledAt: destination.DisabledAt,
		},
		Timestamp: deliveryEvent.Delivery.Time,
	}

	if !attempt.Success && err != nil {
		// Extract attempt data if available
		var delErr *DeliveryError
		if errors.As(err, &delErr) {
			var pubErr *destregistry.ErrDestinationPublishAttempt
			if errors.As(delErr.err, &pubErr) {
				attempt.DeliveryResponse = pubErr.Data
			} else {
				attempt.DeliveryResponse = map[string]interface{}{
					"error": delErr.err.Error(),
				}
			}
		} else {
			attempt.DeliveryResponse = map[string]interface{}{
				"error":   "unexpected",
				"message": err.Error(),
			}
		}
	}

	if monitorErr := h.alertMonitor.HandleAttempt(ctx, attempt); monitorErr != nil {
		h.logger.Ctx(ctx).Error("failed to handle alert attempt",
			zap.Error(monitorErr),
			zap.String("delivery_event_id", deliveryEvent.ID),
			zap.String("destination_id", destination.ID))
		return
	}

	h.logger.Ctx(ctx).Info("alert attempt handled",
		zap.String("delivery_event_id", deliveryEvent.ID),
		zap.String("destination_id", destination.ID))
}

func (h *messageHandler) shouldScheduleRetry(deliveryEvent models.DeliveryEvent, err error) bool {
	if deliveryEvent.Manual {
		return false
	}
	if !deliveryEvent.Event.EligibleForRetry {
		return false
	}
	if _, ok := err.(*destregistry.ErrDestinationPublishAttempt); !ok {
		return false
	}
	// Attempt starts at 0 for initial attempt, so we can compare directly
	return deliveryEvent.Attempt < h.retryMaxLimit
}

func (h *messageHandler) shouldNackError(err error) bool {
	if err == nil {
		return false // Success case, always ack
	}

	// Handle pre-delivery errors (system errors)
	var preErr *PreDeliveryError
	if errors.As(err, &preErr) {
		// Don't nack if it's a permanent error
		if errors.Is(preErr.err, models.ErrDestinationDeleted) || errors.Is(preErr.err, errDestinationDisabled) {
			return false
		}
		return true // Nack other pre-delivery errors
	}

	// Handle delivery errors
	var delErr *DeliveryError
	if errors.As(err, &delErr) {
		return h.shouldNackDeliveryError(delErr.err)
	}

	// Handle post-delivery errors
	var postErr *PostDeliveryError
	if errors.As(err, &postErr) {
		// Check if this wraps a delivery error
		var delErr *DeliveryError
		if errors.As(postErr.err, &delErr) {
			return h.shouldNackDeliveryError(delErr.err)
		}
		return true // Nack other post-delivery errors
	}

	// For any other error type, nack for safety
	return true
}

func (h *messageHandler) shouldNackDeliveryError(err error) bool {
	// Don't nack if it's a delivery attempt error (handled by retry scheduling)
	if _, ok := err.(*destregistry.ErrDestinationPublishAttempt); ok {
		return false
	}
	return true // Nack other delivery errors
}

func (h *messageHandler) scheduleRetry(ctx context.Context, deliveryEvent models.DeliveryEvent) error {
	backoffDuration := h.retryBackoff.Duration(deliveryEvent.Attempt)

	retryMessage := RetryMessageFromDeliveryEvent(deliveryEvent)
	retryMessageStr, err := retryMessage.ToString()
	if err != nil {
		return err
	}

	if err := h.retryScheduler.Schedule(ctx, retryMessageStr, backoffDuration, scheduler.WithTaskID(deliveryEvent.GetRetryID())); err != nil {
		h.logger.Ctx(ctx).Error("failed to schedule retry",
			zap.Error(err),
			zap.String("delivery_event_id", deliveryEvent.ID),
			zap.Int("attempt", deliveryEvent.Attempt),
			zap.Duration("backoff", backoffDuration))
		return err
	}

	h.logger.Ctx(ctx).Audit("retry scheduled",
		zap.String("delivery_event_id", deliveryEvent.ID),
		zap.Int("attempt", deliveryEvent.Attempt),
		zap.Duration("backoff", backoffDuration))

	return nil
}

// ensureDeliveryEvent ensures that the delivery event struct has full data.
// In retry scenarios, the delivery event only has its ID and we'll need to query the full data.
func (h *messageHandler) ensureDeliveryEvent(ctx context.Context, deliveryEvent *models.DeliveryEvent) error {
	// TODO: consider a more deliberate way to check for retry scenario?
	if !deliveryEvent.Event.Time.IsZero() {
		return nil
	}

	event, err := h.logStore.RetrieveEvent(ctx, deliveryEvent.Event.TenantID, deliveryEvent.Event.ID)
	if err != nil {
		return err
	}
	if event == nil {
		return errors.New("event not found")
	}
	deliveryEvent.Event = *event

	return nil
}

// ensurePublishableDestination ensures that the destination exists and is in a publishable state.
// Returns an error if the destination is not found, deleted, disabled, or any other state that
// would prevent publishing.
func (h *messageHandler) ensurePublishableDestination(ctx context.Context, deliveryEvent models.DeliveryEvent) (*models.Destination, error) {
	destination, err := h.entityStore.RetrieveDestination(ctx, deliveryEvent.Event.TenantID, deliveryEvent.DestinationID)
	if err != nil {
		logger := h.logger.Ctx(ctx)
		fields := []zap.Field{
			zap.Error(err),
			zap.String("destination_id", deliveryEvent.DestinationID),
			zap.String("tenant_id", deliveryEvent.Event.TenantID),
		}

		if errors.Is(err, models.ErrDestinationDeleted) {
			logger.Info("destination deleted", fields...)
		} else {
			// Unexpected errors like DB connection issues
			logger.Error("failed to retrieve destination", fields...)
		}
		return nil, err
	}
	if destination == nil {
		h.logger.Ctx(ctx).Info("destination not found",
			zap.String("destination_id", deliveryEvent.DestinationID),
			zap.String("tenant_id", deliveryEvent.Event.TenantID))
		return nil, models.ErrDestinationNotFound
	}
	if destination.DisabledAt != nil {
		h.logger.Ctx(ctx).Info("skipping disabled destination",
			zap.String("destination_id", destination.ID),
			zap.String("destination_type", destination.Type),
			zap.Time("disabled_at", *destination.DisabledAt))
		return nil, errDestinationDisabled
	}
	return destination, nil
}
