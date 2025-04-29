# DestinationUpdate


## Supported Types

### `components.DestinationUpdateWebhook`

```typescript
const value: components.DestinationUpdateWebhook = {
  topics: "*",
  config: {
    url: "https://example.com/webhooks/user",
  },
};
```

### `components.DestinationUpdateAWSSQS`

```typescript
const value: components.DestinationUpdateAWSSQS = {
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

### `components.DestinationUpdateRabbitMQ`

```typescript
const value: components.DestinationUpdateRabbitMQ = {
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

