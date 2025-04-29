# DestinationUpdateAWSSQS

## Example Usage

```typescript
import { DestinationUpdateAWSSQS } from "openapi/models/components";

let value: DestinationUpdateAWSSQS = {
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

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `topics`                                                                     | *components.Topics*                                                          | :heavy_minus_sign:                                                           | "*" or an array of enabled topics.                                           | *                                                                            |
| `config`                                                                     | [components.AWSSQSConfig](../../models/components/awssqsconfig.md)           | :heavy_minus_sign:                                                           | N/A                                                                          |                                                                              |
| `credentials`                                                                | [components.AWSSQSCredentials](../../models/components/awssqscredentials.md) | :heavy_minus_sign:                                                           | N/A                                                                          |                                                                              |