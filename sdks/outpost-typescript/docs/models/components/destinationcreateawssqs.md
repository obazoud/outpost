# DestinationCreateAWSSQS

## Example Usage

```typescript
import { DestinationCreateAWSSQS } from "@hookdeck/outpost-sdk/models/components";

let value: DestinationCreateAWSSQS = {
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

## Fields

| Field                                                                                            | Type                                                                                             | Required                                                                                         | Description                                                                                      | Example                                                                                          |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `id`                                                                                             | *string*                                                                                         | :heavy_minus_sign:                                                                               | Optional user-provided ID. A UUID will be generated if empty.                                    | user-provided-id                                                                                 |
| `type`                                                                                           | [components.DestinationCreateAWSSQSType](../../models/components/destinationcreateawssqstype.md) | :heavy_check_mark:                                                                               | Type of the destination. Must be 'aws_sqs'.                                                      |                                                                                                  |
| `topics`                                                                                         | *components.Topics*                                                                              | :heavy_check_mark:                                                                               | "*" or an array of enabled topics.                                                               | *                                                                                                |
| `config`                                                                                         | [components.AWSSQSConfig](../../models/components/awssqsconfig.md)                               | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |
| `credentials`                                                                                    | [components.AWSSQSCredentials](../../models/components/awssqscredentials.md)                     | :heavy_check_mark:                                                                               | N/A                                                                                              |                                                                                                  |