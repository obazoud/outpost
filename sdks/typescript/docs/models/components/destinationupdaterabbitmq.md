# DestinationUpdateRabbitMQ

## Example Usage

```typescript
import { DestinationUpdateRabbitMQ } from "openapi/models/components";

let value: DestinationUpdateRabbitMQ = {
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

## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      | Example                                                                          |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `topics`                                                                         | *components.Topics*                                                              | :heavy_minus_sign:                                                               | "*" or an array of enabled topics.                                               | *                                                                                |
| `config`                                                                         | [components.RabbitMQConfig](../../models/components/rabbitmqconfig.md)           | :heavy_minus_sign:                                                               | N/A                                                                              |                                                                                  |
| `credentials`                                                                    | [components.RabbitMQCredentials](../../models/components/rabbitmqcredentials.md) | :heavy_minus_sign:                                                               | N/A                                                                              |                                                                                  |