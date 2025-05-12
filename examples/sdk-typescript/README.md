## Outpost TypeScript SDK Example

This example demonstrates using the Outpost TypeScript SDK.

### Prerequisites

*   Node.js (LTS recommended)
*   NPM

> [!NOTE]
> All commands below should be run from within the `examples/sdk-typescript` directory.

### Setup

1.  **Install dependencies:**
    ```bash
    npm install
    ```

### Running the Example

1.  **Configure environment variables:**
    Create a `.env` file in this directory (`examples/sdk-typescript`) with the following:
    ```dotenv
    SERVER_URL="your_server_url"
    ADMIN_API_KEY="your_admin_api_key"
    TENANT_ID="your_tenant_id"
    ```
    Replace the placeholder values with your Outpost server URL, Admin API key, and Tenant ID. (Note: `.env` is already gitignored).

2.  **Run the example:**
    ```bash
    npm run start
    ```

    This executes `index.ts` via `ts-node`, showcasing SDK functionalities. Review `index.ts` for details.
