# DestinationCreate


## Supported Types

### `components.DestinationCreateWebhook`

```typescript
const value: components.DestinationCreateWebhook = {
  id: "user-provided-id",
  type: "webhook",
  topics: "*",
  config: {
    url: "https://example.com/webhooks/user",
  },
  credentials: {
    secret: "whsec_abc123",
    previousSecret: "whsec_xyz789",
    previousSecretInvalidAt: new Date("2024-01-02T00:00:00Z"),
  },
};
```

### `components.DestinationCreateAWSSQS`

```typescript
const value: components.DestinationCreateAWSSQS = {
  id: "user-provided-id",
  type: "aws_sqs",
  topics: "*",
  config: {
    endpoint: "https://sqs.us-east-1.amazonaws.com",
    queueUrl: "https://sqs.us-east-1.amazonaws.com/123456789012/my-queue",
  },
  credentials: {
    key: "AKIAIOSFODNN7EXAMPLE",
    secret: "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
    session: "AQoDYXdzEPT//////////wEXAMPLE...",
  },
};
```

### `components.DestinationCreateRabbitMQ`

```typescript
const value: components.DestinationCreateRabbitMQ = {
  id: "user-provided-id",
  type: "rabbitmq",
  topics: "*",
  config: {
    serverUrl: "localhost:5672",
    exchange: "my-exchange",
    tls: "false",
  },
  credentials: {
    username: "guest",
    password: "guest",
  },
};
```

