# DestinationOAuthRedirect

## Example Usage

```typescript
import { DestinationOAuthRedirect } from "openapi/models/components";

let value: DestinationOAuthRedirect = {
  redirectUrl: "https://dashboard.hookdeck.com/authorize?token=12313123",
};
```

## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `redirectUrl`                                                        | *string*                                                             | :heavy_minus_sign:                                                   | Redirect URL for OAuth flow if applicable during destination update. | https://dashboard.hookdeck.com/authorize?token=12313123              |