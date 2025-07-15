# AzureServiceBus Destination configuration

Here's a rough document explaining how AzureServiceBus works and how the destination is implemented with Outpost.

# PubSub vs Queue

Azure ServiceBus supports both PubSub (Topic & Subscription) and Queue. From the Publisher (Azure's term is Sender) perspective, it doesn't really care whether it's publishing to a Topic or to a Queue. So, from the destination config, all we need is a single "name" field.

## Authentication

For authentication, we currently support "connection_string" which generally have access to the full Namespace. So if the end-user wants to ensure Outpost only has access to their desired queue or topic, they should create a new Namespace just for Outpost.

## Message

Whether it's publishing to Topic or Queue, the Publisher needs to send an Azure's Message. Here's the full Golang SDK Message struct:

```golang
// Message is a message with a body and commonly used properties.
// Properties that are pointers are optional.
type Message struct {
	// ApplicationProperties can be used to store custom metadata for a message.
	ApplicationProperties map[string]any

	// Body corresponds to the first []byte array in the Data section of an AMQP message.
	Body []byte

	// ContentType describes the payload of the message, with a descriptor following
	// the format of Content-Type, specified by RFC2045 (ex: "application/json").
	ContentType *string

	// CorrelationID allows an application to specify a context for the message for the purposes of
	// correlation, for example reflecting the MessageID of a message that is being
	// replied to.
	CorrelationID *string

	// MessageID is an application-defined value that uniquely identifies
	// the message and its payload. The identifier is a free-form string.
	//
	// If enabled, the duplicate detection feature identifies and removes further submissions
	// of messages with the same MessageId.
	MessageID *string

	// PartitionKey is used with a partitioned entity and enables assigning related messages
	// to the same internal partition. This ensures that the submission sequence order is correctly
	// recorded. The partition is chosen by a hash function in Service Bus and cannot be chosen
	// directly.
	//
	// For session-aware entities, the ReceivedMessage.SessionID overrides this value.
	PartitionKey *string

	// ReplyTo is an application-defined value specify a reply path to the receiver of the message. When
	// a sender expects a reply, it sets the value to the absolute or relative path of the queue or topic
	// it expects the reply to be sent to.
	ReplyTo *string

	// ReplyToSessionID augments the ReplyTo information and specifies which SessionId should
	// be set for the reply when sent to the reply entity.
	ReplyToSessionID *string

	// ScheduledEnqueueTime specifies a time when a message will be enqueued. The message is transferred
	// to the broker but will not available until the scheduled time.
	ScheduledEnqueueTime *time.Time

	// SessionID is used with session-aware entities and associates a message with an application-defined
	// session ID. Note that an empty string is a valid session identifier.
	// Messages with the same session identifier are subject to summary locking and enable
	// exact in-order processing and demultiplexing. For session-unaware entities, this value is ignored.
	SessionID *string

	// Subject enables an application to indicate the purpose of the message, similar to an email subject line.
	Subject *string

	// TimeToLive is the duration after which the message expires, starting from the instant the
	// message has been accepted and stored by the broker, found in the ReceivedMessage.EnqueuedTime
	// property.
	//
	// When not set explicitly, the assumed value is the DefaultTimeToLive for the queue or topic.
	// A message's TimeToLive cannot be longer than the entity's DefaultTimeToLive is silently
	// adjusted if it does.
	TimeToLive *time.Duration

	// To is reserved for future use in routing scenarios but is not currently used by Service Bus.
	// Applications can use this value to indicate the logical destination of the message.
	To *string
}
```

Here are a few notable configuration, especially on the destination level that we may want to support:

- MessageID --> MessageIDTemplate, similar to AWS Kinesis Parition Key approach
- CorrelationID --> CorrelationIDTemplate, similar to AWS Kinesis Parition Key approach
- PartitionKey --> PartitionKeyTemplate, similar to AWS Kinesis Parition Key approach

- ScheduledEnqueueTime
- TimeToLive

The current implementation doesn't support any of these. So when create destination, it's super straightforward:

```golang
type Config struct {
  Name string
}
```

If we want to support these, we can either add them to Config, such as `Config.TTL`, or we can also add a suffix like `Config.MessageTTL` to specify that these config would apply to the Message.
