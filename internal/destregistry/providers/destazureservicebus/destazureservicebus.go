package destazureservicebus

import (
	"context"
	"fmt"

	"github.com/hookdeck/outpost/internal/destregistry"
	"github.com/hookdeck/outpost/internal/destregistry/metadata"
	"github.com/hookdeck/outpost/internal/models"
)

type AzureServiceBusDestination struct {
	*destregistry.BaseProvider
}

type AzureServiceBusDestinationConfig struct {
	Topic string
}

type AzureServiceBusDestinationCredentials struct {
	ConnectionString string
}

var _ destregistry.Provider = (*AzureServiceBusDestination)(nil)

func New(loader metadata.MetadataLoader) (*AzureServiceBusDestination, error) {
	base, err := destregistry.NewBaseProvider(loader, "azure_servicebus")
	if err != nil {
		return nil, err
	}

	return &AzureServiceBusDestination{
		BaseProvider: base,
	}, nil
}

func (d *AzureServiceBusDestination) Validate(ctx context.Context, destination *models.Destination) error {
	// For phase 1, just call base validation
	return d.BaseProvider.Validate(ctx, destination)
}

func (d *AzureServiceBusDestination) CreatePublisher(ctx context.Context, destination *models.Destination) (destregistry.Publisher, error) {
	cfg, creds, err := d.resolveMetadata(ctx, destination)
	if err != nil {
		return nil, err
	}

	return &AzureServiceBusPublisher{
		BasePublisher:    &destregistry.BasePublisher{},
		connectionString: creds.ConnectionString,
		topic:            cfg.Topic,
	}, nil
}

func (d *AzureServiceBusDestination) ComputeTarget(destination *models.Destination) destregistry.DestinationTarget {
	if topic, ok := destination.Config["topic"]; ok {
		return destregistry.DestinationTarget{
			Target:    topic,
			TargetURL: "",
		}
	}
	return destregistry.DestinationTarget{}
}

func (d *AzureServiceBusDestination) Preprocess(newDestination *models.Destination, originalDestination *models.Destination, opts *destregistry.PreprocessDestinationOpts) error {
	// Phase 1: empty implementation
	return nil
}

func (d *AzureServiceBusDestination) resolveMetadata(ctx context.Context, destination *models.Destination) (*AzureServiceBusDestinationConfig, *AzureServiceBusDestinationCredentials, error) {
	if err := d.BaseProvider.Validate(ctx, destination); err != nil {
		return nil, nil, err
	}

	return &AzureServiceBusDestinationConfig{
			Topic: destination.Config["topic"],
		}, &AzureServiceBusDestinationCredentials{
			ConnectionString: destination.Credentials["connection_string"],
		}, nil
}

// Publisher implementation
type AzureServiceBusPublisher struct {
	*destregistry.BasePublisher
	connectionString string
	topic            string
}

func (p *AzureServiceBusPublisher) Publish(ctx context.Context, event *models.Event) (*destregistry.Delivery, error) {
	// Phase 1: minimal implementation that returns an error
	return nil, fmt.Errorf("Azure Service Bus publishing not yet implemented")
}

func (p *AzureServiceBusPublisher) Close() error {
	p.BasePublisher.StartClose()
	// Phase 1: empty implementation
	return nil
}
