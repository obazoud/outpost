# DestinationCreateRabbitMQ

## Example Usage

```typescript
import { DestinationCreateRabbitMQ } from "openapi/models/components";

let value: DestinationCreateRabbitMQ = {
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

## Fields

| Field                                                                                                | Type                                                                                                 | Required                                                                                             | Description                                                                                          | Example                                                                                              |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `id`                                                                                                 | *string*                                                                                             | :heavy_minus_sign:                                                                                   | Optional user-provided ID. A UUID will be generated if empty.                                        | user-provided-id                                                                                     |
| `type`                                                                                               | [components.DestinationCreateRabbitMQType](../../models/components/destinationcreaterabbitmqtype.md) | :heavy_check_mark:                                                                                   | Type of the destination. Must be 'rabbitmq'.                                                         |                                                                                                      |
| `topics`                                                                                             | *components.Topics*                                                                                  | :heavy_check_mark:                                                                                   | "*" or an array of enabled topics.                                                                   | *                                                                                                    |
| `config`                                                                                             | [components.RabbitMQConfig](../../models/components/rabbitmqconfig.md)                               | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |
| `credentials`                                                                                        | [components.RabbitMQCredentials](../../models/components/rabbitmqcredentials.md)                     | :heavy_check_mark:                                                                                   | N/A                                                                                                  |                                                                                                      |