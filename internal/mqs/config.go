package mqs

type QueueConfig struct {
	AWSSQS          *AWSSQSConfig
	AzureServiceBus *AzureServiceBusConfig
	GCPPubSub       *GCPPubSubConfig
	RabbitMQ        *RabbitMQConfig
	InMemory        *InMemoryConfig // mainly for testing purposes
}

type AzureServiceBusConfig struct {
}

type GCPPubSubConfig struct {
	ProjectID                 string
	TopicID                   string
	SubscriptionID            string
	ServiceAccountCredentials string // JSON key file content
}

type InMemoryConfig struct {
	Name string
}
