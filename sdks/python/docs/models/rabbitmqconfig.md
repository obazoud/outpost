# RabbitMQConfig


## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `server_url`                                                | *str*                                                       | :heavy_check_mark:                                          | RabbitMQ server address (host:port).                        | localhost:5672                                              |
| `exchange`                                                  | *str*                                                       | :heavy_check_mark:                                          | The exchange to publish messages to.                        | my-exchange                                                 |
| `tls`                                                       | [Optional[models.TLS]](../models/tls.md)                    | :heavy_minus_sign:                                          | Whether to use TLS connection (amqps). Defaults to "false". | false                                                       |