# RabbitMQConfig

## Example Usage

```typescript
import { RabbitMQConfig } from "openapi/models/components";

let value: RabbitMQConfig = {
  serverUrl: "localhost:5672",
  exchange: "my-exchange",
  tls: "false",
};
```

## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `serverUrl`                                                 | *string*                                                    | :heavy_check_mark:                                          | RabbitMQ server address (host:port).                        | localhost:5672                                              |
| `exchange`                                                  | *string*                                                    | :heavy_check_mark:                                          | The exchange to publish messages to.                        | my-exchange                                                 |
| `tls`                                                       | [components.Tls](../../models/components/tls.md)            | :heavy_minus_sign:                                          | Whether to use TLS connection (amqps). Defaults to "false". | false                                                       |