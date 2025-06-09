package config

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicebus/armservicebus"
	"github.com/hookdeck/outpost/internal/mqinfra"
	"github.com/hookdeck/outpost/internal/mqs"
)

type AzureServiceBusConfig struct {
	TenantID       string `yaml:"tenant_id" env:"AZURE_SERVICEBUS_TENANT_ID" desc:"Azure Active Directory tenant ID" required:"Y"`
	ClientID       string `yaml:"client_id" env:"AZURE_SERVICEBUS_CLIENT_ID" desc:"Service principal client ID" required:"Y"`
	ClientSecret   string `yaml:"client_secret" env:"AZURE_SERVICEBUS_CLIENT_SECRET" desc:"Service principal client secret" required:"Y"`
	SubscriptionID string `yaml:"subscription_id" env:"AZURE_SERVICEBUS_SUBSCRIPTION_ID" desc:"Azure subscription ID" required:"Y"`
	ResourceGroup  string `yaml:"resource_group" env:"AZURE_SERVICEBUS_RESOURCE_GROUP" desc:"Azure resource group name" required:"Y"`
	Namespace      string `yaml:"namespace" env:"AZURE_SERVICEBUS_NAMESPACE" desc:"Azure Service Bus namespace" required:"Y"`

	DeliveryTopic        string `yaml:"delivery_topic" env:"AZURE_SERVICEBUS_DELIVERY_TOPIC" desc:"Topic name for delivery queue" required:"N" default:"outpost-delivery"`
	DeliverySubscription string `yaml:"delivery_subscription" env:"AZURE_SERVICEBUS_DELIVERY_SUBSCRIPTION" desc:"Subscription name for delivery queue" required:"N" default:"outpost-delivery-subscription"`
	LogTopic             string `yaml:"log_topic" env:"AZURE_SERVICEBUS_LOG_TOPIC" desc:"Topic name for log queue" required:"N" default:"outpost-log"`
	LogSubscription      string `yaml:"log_subscription" env:"AZURE_SERVICEBUS_LOG_SUBSCRIPTION" desc:"Subscription name for log queue" required:"N" default:"outpost-log-subscription"`

	// connectionStringOnce  sync.Once
	// connectionString      string
	// connectionStringError error
}

func (c *AzureServiceBusConfig) IsConfigured() bool {
	return c.TenantID != "" && c.ClientID != "" && c.ClientSecret != "" && c.SubscriptionID != "" && c.ResourceGroup != "" && c.Namespace != ""
}

func (c *AzureServiceBusConfig) GetProviderType() string {
	return "azure_service_bus"
}

func (c *AzureServiceBusConfig) getTopicByQueueType(queueType string) string {
	switch queueType {
	case "deliverymq":
		return c.DeliveryTopic
	case "logmq":
		return c.LogTopic
	default:
		return ""
	}
}

func (c *AzureServiceBusConfig) getSubscriptionByQueueType(queueType string) string {
	switch queueType {
	case "deliverymq":
		return c.DeliverySubscription
	case "logmq":
		return c.LogSubscription
	default:
		return ""
	}
}

func (c *AzureServiceBusConfig) ToInfraConfig(queueType string) *mqinfra.MQInfraConfig {
	if !c.IsConfigured() {
		return nil
	}

	topic := c.getTopicByQueueType(queueType)
	subscription := c.getSubscriptionByQueueType(queueType)

	return &mqinfra.MQInfraConfig{
		AzureServiceBus: &mqinfra.AzureServiceBusInfraConfig{
			TenantID:       c.TenantID,
			ClientID:       c.ClientID,
			ClientSecret:   c.ClientSecret,
			SubscriptionID: c.SubscriptionID,
			ResourceGroup:  c.ResourceGroup,
			Namespace:      c.Namespace,
			Topic:          topic,
			Subscription:   subscription,
		},
	}
}

func (c *AzureServiceBusConfig) ToQueueConfig(ctx context.Context, queueType string) (*mqs.QueueConfig, error) {
	if !c.IsConfigured() {
		return nil, nil
	}

	// connectionString, err := c.getConnectionString(ctx)
	// if err != nil {
	// 	return nil, err
	// }

	topic := c.getTopicByQueueType(queueType)
	subscription := c.getSubscriptionByQueueType(queueType)

	return &mqs.QueueConfig{
		AzureServiceBus: &mqs.AzureServiceBusConfig{
			Topic:          topic,
			Subscription:   subscription,
			TenantID:       c.TenantID,
			ClientID:       c.ClientID,
			ClientSecret:   c.ClientSecret,
			SubscriptionID: c.SubscriptionID,
			ResourceGroup:  c.ResourceGroup,
			Namespace:      c.Namespace,
		},
	}, nil
}

func (c *AzureServiceBusConfig) getConnectionString(ctx context.Context) (string, error) {
	c.connectionStringOnce.Do(func() {
		cred, err := azidentity.NewClientSecretCredential(
			c.TenantID,
			c.ClientID,
			c.ClientSecret,
			nil,
		)
		if err != nil {
			c.connectionStringError = fmt.Errorf("failed to create credential: %w", err)
			return
		}

		sbClient, err := armservicebus.NewNamespacesClient(c.SubscriptionID, cred, nil)
		if err != nil {
			c.connectionStringError = fmt.Errorf("failed to create servicebus client: %w", err)
			return
		}

		keysResp, err := sbClient.ListKeys(ctx, c.ResourceGroup, c.Namespace, "RootManageSharedAccessKey", nil)
		if err != nil {
			c.connectionStringError = fmt.Errorf("failed to get keys: %w", err)
			return
		}

		if keysResp.PrimaryConnectionString == nil {
			c.connectionStringError = fmt.Errorf("no connection string found")
			return
		}

		c.connectionString = *keysResp.PrimaryConnectionString
	})

	return c.connectionString, c.connectionStringError
}
