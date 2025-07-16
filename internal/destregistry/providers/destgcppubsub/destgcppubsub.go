package destgcppubsub

import (
	"context"
	"fmt"

	"github.com/hookdeck/outpost/internal/destregistry"
	"github.com/hookdeck/outpost/internal/destregistry/metadata"
	"github.com/hookdeck/outpost/internal/models"
)

type GCPPubSubDestination struct {
	*destregistry.BaseProvider
}

type GCPPubSubDestinationConfig struct {
	ProjectID string
	TopicName string
}

type GCPPubSubDestinationCredentials struct {
	ServiceAccountJSON string
}

var _ destregistry.Provider = (*GCPPubSubDestination)(nil)

func New(loader metadata.MetadataLoader) (*GCPPubSubDestination, error) {
	base, err := destregistry.NewBaseProvider(loader, "gcp_pubsub")
	if err != nil {
		return nil, err
	}

	return &GCPPubSubDestination{
		BaseProvider: base,
	}, nil
}

func (d *GCPPubSubDestination) Validate(ctx context.Context, destination *models.Destination) error {
	_, _, err := d.resolveMetadata(ctx, destination)
	if err != nil {
		return err
	}
	return nil
}

func (d *GCPPubSubDestination) CreatePublisher(ctx context.Context, destination *models.Destination) (destregistry.Publisher, error) {
	// TODO: Implement
	return nil, fmt.Errorf("not implemented")
}

func (d *GCPPubSubDestination) resolveMetadata(ctx context.Context, destination *models.Destination) (*GCPPubSubDestinationConfig, *GCPPubSubDestinationCredentials, error) {
	if err := d.BaseProvider.Validate(ctx, destination); err != nil {
		return nil, nil, err
	}

	return &GCPPubSubDestinationConfig{
			ProjectID: destination.Config["project_id"],
			TopicName: destination.Config["topic_name"],
		}, &GCPPubSubDestinationCredentials{
			ServiceAccountJSON: destination.Credentials["service_account_json"],
		}, nil
}

func (d *GCPPubSubDestination) ComputeTarget(destination *models.Destination) destregistry.DestinationTarget {
	// TODO: Return a human-readable target description
	return destregistry.DestinationTarget{
		Target:    "unknown",
		TargetURL: "", // Optional: URL to view the resource in a web console
	}
}

type GCPPubSubPublisher struct {
	*destregistry.BasePublisher
	
	// TODO: Add publisher fields
}

func (pub *GCPPubSubPublisher) Publish(ctx context.Context, event *models.Event) (*destregistry.Delivery, error) {
	if err := pub.BasePublisher.StartPublish(); err != nil {
		return nil, err
	}
	defer pub.BasePublisher.FinishPublish()

	// TODO: Implement
	return nil, fmt.Errorf("not implemented")
}

func (pub *GCPPubSubPublisher) Close() error {
	pub.BasePublisher.StartClose()

	// TODO: Add cleanup
	return nil
}