# AWSSQSConfig

## Example Usage

```typescript
import { AWSSQSConfig } from "openapi/models/components";

let value: AWSSQSConfig = {
  endpoint: "https://sqs.us-east-1.amazonaws.com",
  queueUrl: "https://sqs.us-east-1.amazonaws.com/123456789012/my-queue",
};
```

## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   | Example                                                                       |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `endpoint`                                                                    | *string*                                                                      | :heavy_minus_sign:                                                            | Optional. Custom AWS endpoint URL (e.g., for LocalStack or specific regions). | https://sqs.us-east-1.amazonaws.com                                           |
| `queueUrl`                                                                    | *string*                                                                      | :heavy_check_mark:                                                            | The URL of the SQS queue.                                                     | https://sqs.us-east-1.amazonaws.com/123456789012/my-queue                     |