package config

import (
	"fmt"

	"github.com/hookdeck/outpost/internal/mqs"
)

// MQ Infrastructure configs
type AWSSQSConfig struct {
	AccessKeyID     string `yaml:"access_key_id" env:"AWS_SQS_ACCESS_KEY_ID" desc:"AWS Access Key ID for SQS. Required if AWS SQS is the chosen MQ provider." required:"C"`
	SecretAccessKey string `yaml:"secret_access_key" env:"AWS_SQS_SECRET_ACCESS_KEY" desc:"AWS Secret Access Key for SQS. Required if AWS SQS is the chosen MQ provider." required:"C"`
	Region          string `yaml:"region" env:"AWS_SQS_REGION" desc:"AWS Region for SQS. Required if AWS SQS is the chosen MQ provider." required:"C"`
	Endpoint        string `yaml:"endpoint" env:"AWS_SQS_ENDPOINT" desc:"Custom AWS SQS endpoint URL. Optional, typically used for local testing (e.g., LocalStack)." required:"N"`
	DeliveryQueue   string `yaml:"delivery_queue" env:"AWS_SQS_DELIVERY_QUEUE" desc:"Name of the SQS queue for delivery events." required:"N"`
	LogQueue        string `yaml:"log_queue" env:"AWS_SQS_LOG_QUEUE" desc:"Name of the SQS queue for log events." required:"N"`
}

type GCPPubSubConfig struct {
	Project                   string `yaml:"project" env:"GCP_PUBSUB_PROJECT" desc:"GCP Project ID for Pub/Sub. Required if GCP Pub/Sub is the chosen MQ provider." required:"C"`
	ServiceAccountCredentials string `yaml:"service_account_credentials" env:"GCP_PUBSUB_SERVICE_ACCOUNT_CREDENTIALS" desc:"JSON string or path to a file containing GCP service account credentials for Pub/Sub. Required if GCP Pub/Sub is the chosen MQ provider and not running in an environment with implicit credentials (e.g., GCE, GKE)." required:"C"`
	DeliveryTopic             string `yaml:"delivery_topic" env:"GCP_PUBSUB_DELIVERY_TOPIC" desc:"Name of the GCP Pub/Sub topic for delivery events." required:"N"`
	DeliverySubscription      string `yaml:"delivery_subscription" env:"GCP_PUBSUB_DELIVERY_SUBSCRIPTION" desc:"Name of the GCP Pub/Sub subscription for delivery events." required:"N"`
	LogTopic                  string `yaml:"log_topic" env:"GCP_PUBSUB_LOG_TOPIC" desc:"Name of the GCP Pub/Sub topic for log events." required:"N"`
	LogSubscription           string `yaml:"log_subscription" env:"GCP_PUBSUB_LOG_SUBSCRIPTION" desc:"Name of the GCP Pub/Sub subscription for log events." required:"N"`
}

type RabbitMQConfig struct {
	ServerURL     string `yaml:"server_url" env:"RABBITMQ_SERVER_URL" desc:"RabbitMQ server connection URL (e.g., 'amqp://user:pass@host:port/vhost'). Required if RabbitMQ is the chosen MQ provider." required:"C"`
	Exchange      string `yaml:"exchange" env:"RABBITMQ_EXCHANGE" desc:"Name of the RabbitMQ exchange to use." required:"N"`
	DeliveryQueue string `yaml:"delivery_queue" env:"RABBITMQ_DELIVERY_QUEUE" desc:"Name of the RabbitMQ queue for delivery events." required:"N"`
	LogQueue      string `yaml:"log_queue" env:"RABBITMQ_LOG_QUEUE" desc:"Name of the RabbitMQ queue for log events." required:"N"`
}

type MQsConfig struct {
	AWSSQS    AWSSQSConfig    `yaml:"aws_sqs" desc:"Configuration for using AWS SQS as the message queue. Only one MQ provider (AWSSQS, GCPPubSub, RabbitMQ) should be configured." required:"N"`
	GCPPubSub GCPPubSubConfig `yaml:"gcp_pubsub" desc:"Configuration for using GCP Pub/Sub as the message queue. Only one MQ provider (AWSSQS, GCPPubSub, RabbitMQ) should be configured." required:"N"`
	RabbitMQ  RabbitMQConfig  `yaml:"rabbitmq" desc:"Configuration for using RabbitMQ as the message queue. Only one MQ provider (AWSSQS, GCPPubSub, RabbitMQ) should be configured." required:"N"`
}

func (c MQsConfig) GetInfraType() string {
	if hasAWSSQSConfig(c.AWSSQS) {
		return "awssqs"
	}
	if hasGCPPubSubConfig(c.GCPPubSub) {
		return "gcppubsub"
	}
	if hasRabbitMQConfig(c.RabbitMQ) {
		return "rabbitmq"
	}
	return ""
}

// getQueueConfig returns a queue config for the given queue type
// queueType can be "deliverymq" or "logmq"
func (c *MQsConfig) getQueueConfig(queueType string) *mqs.QueueConfig {
	if c == nil {
		return nil
	}

	infraType := c.GetInfraType()
	switch infraType {
	case "awssqs":
		queue := ""
		if queueType == "deliverymq" {
			queue = c.AWSSQS.DeliveryQueue
		} else if queueType == "logmq" {
			queue = c.AWSSQS.LogQueue
		}

		creds := fmt.Sprintf("%s:%s:", c.AWSSQS.AccessKeyID, c.AWSSQS.SecretAccessKey)
		return &mqs.QueueConfig{
			AWSSQS: &mqs.AWSSQSConfig{
				Endpoint:                  c.AWSSQS.Endpoint,
				Region:                    c.AWSSQS.Region,
				ServiceAccountCredentials: creds,
				Topic:                     queue,
			},
		}
	case "gcppubsub":
		topic := ""
		subscription := ""
		if queueType == "deliverymq" {
			topic = c.GCPPubSub.DeliveryTopic
			subscription = c.GCPPubSub.DeliverySubscription
		} else if queueType == "logmq" {
			topic = c.GCPPubSub.LogTopic
			subscription = c.GCPPubSub.LogSubscription
		}
		return &mqs.QueueConfig{
			GCPPubSub: &mqs.GCPPubSubConfig{
				ProjectID:                 c.GCPPubSub.Project,
				ServiceAccountCredentials: c.GCPPubSub.ServiceAccountCredentials,
				TopicID:                   topic,
				SubscriptionID:            subscription,
			},
		}
	case "rabbitmq":
		queue := ""
		if queueType == "deliverymq" {
			queue = c.RabbitMQ.DeliveryQueue
		} else if queueType == "logmq" {
			queue = c.RabbitMQ.LogQueue
		}
		return &mqs.QueueConfig{
			RabbitMQ: &mqs.RabbitMQConfig{
				ServerURL: c.RabbitMQ.ServerURL,
				Exchange:  c.RabbitMQ.Exchange,
				Queue:     queue,
			},
		}
	default:
		return nil
	}
}

func (c MQsConfig) GetDeliveryQueueConfig() *mqs.QueueConfig {
	infraType := c.GetInfraType()
	switch infraType {
	case "awssqs":
		return c.getQueueConfig("deliverymq")
	case "gcppubsub":
		return c.getQueueConfig("deliverymq")
	case "rabbitmq":
		return c.getQueueConfig("deliverymq")
	default:
		return nil
	}
}

func (c MQsConfig) GetLogQueueConfig() *mqs.QueueConfig {
	infraType := c.GetInfraType()
	switch infraType {
	case "awssqs":
		return c.getQueueConfig("logmq")
	case "gcppubsub":
		return c.getQueueConfig("logmq")
	case "rabbitmq":
		return c.getQueueConfig("logmq")
	default:
		return nil
	}
}

// Helper functions to check for required fields
func hasAWSSQSConfig(config AWSSQSConfig) bool {
	return config.AccessKeyID != "" &&
		config.SecretAccessKey != "" && config.Region != ""
}

func hasGCPPubSubConfig(config GCPPubSubConfig) bool {
	return config.Project != ""
}

func hasRabbitMQConfig(config RabbitMQConfig) bool {
	return config.ServerURL != ""
}
