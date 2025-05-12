## Outpost Go SDK Example

This example demonstrates using the Outpost Go SDK.

The source code for the Go SDK can be found in the [`sdks/outpost-go/`](../../sdks/outpost-go/) directory.

### Prerequisites

*   Go (version specified in `go.mod`)

> [!NOTE]
> All commands below should be run from within the `examples/sdk-go` directory.

### Setup

1.  **Download dependencies:**
    *(This will also add the `godotenv` package needed for `.env` file loading)*
    ```bash
    go mod tidy
    ```

### Running the Example

1.  **Configure environment variables:**
    Create a `.env` file in this directory (`examples/sdk-go`) with the following:
    ```dotenv
    SERVER_URL="your_server_url"
    ADMIN_API_KEY="your_admin_api_key"
    TENANT_ID="your_tenant_id"
    ```
    Replace the placeholder values with your Outpost server URL, Admin API key, and Tenant ID. (Note: You'll need to create a `.gitignore` file if you don't want to commit `.env`).

2.  **Run the example:**
    ```bash
    go run main.go
    ```

    This executes `main.go`, which loads configuration from `.env` and performs a health check using the SDK. Review `main.go` for details.