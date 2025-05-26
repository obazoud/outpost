# Plan to Synchronize Configuration Code and Documentation

**Date:** 2025-05-26

**Objective:** To ensure that the Go configuration code is the single source of truth for all configuration parameters (YAML key, environment variable, description, default value, and required status), and to automatically generate the Markdown documentation (`docs/pages/references/configuration.mdx`) from this source.

## Phase 1: Augment Go Configuration Structs

This phase involves modifying Go source files within the `internal/config/` directory.

1.  **Introduce/Update Struct Tags**:
    *   `desc:"..."`: For the human-readable description of the configuration field. If the field is conditionally required, this description **MUST** include the conditions.
    *   `required:"(Y|N|C)"`:
        *   `Y`: The field is unconditionally required.
        *   `N`: The field is not required.
        *   `C`: The field is conditionally required. The specific conditions **must be detailed in the `desc` tag**.

2.  **Update Struct Fields**:
    Go through all fields in all configuration structs (e.g., `Config`, `RedisConfig`, `MQsConfig`, `DestinationWebhookConfig`, etc., located in `internal/config/`) and add or update these tags.

    *Example Field Definition:*
    ```go
    // In internal/config/config.go
    type Config struct {
        // ...
        LogLevel string `yaml:"log_level" env:"LOG_LEVEL" desc:"Defines the verbosity of application logs (e.g., 'info', 'debug', 'error')." required:"Y"`
        APIKey   string `yaml:"api_key" env:"API_KEY" desc:"The API key for securing the Outpost API. Required only if the API service is enabled." required:"C"`
        // ...
    }

    // In internal/config/redis.go (or wherever RedisConfig is defined)
    type RedisConfig struct {
        Host     string `yaml:"host" env:"REDIS_HOST" desc:"Hostname or IP address of the Redis server." required:"Y"`
        Port     int    `yaml:"port" env:"REDIS_PORT" desc:"Port number for the Redis server." required:"Y"`
        Password string `yaml:"password" env:"REDIS_PASSWORD" desc:"Password for Redis authentication." required:"N"`
        // ...
    }
    ```

## Phase 2: Develop a Documentation Generation Tool (`configdocsgen`)

This phase involves creating a new Go command-line program, likely residing in `cmd/configdocsgen` or `tools/configdocsgen`.

**High-Level Tool Workflow:**

```mermaid
graph TD
    A[Start configdocsgen] --> B{Parse all *.go files in internal/config/};
    B --> C{Identify all config structs};
    C --> D{For each struct field, extract: <br/> - Go field name <br/> - yaml tag <br/> - env tag <br/> - desc tag <br/> - required tag <br/> - Data type};
    D --> E{Instantiate a main Config object};
    E --> F{Call InitDefaults() on the object};
    F --> G{Use reflection to get default values for all fields from the initialized object};
    G --> H{Collect all extracted information into a structured list};
    H --> I{Generate Environment Variables Markdown Table};
    I --> J{Generate YAML Structure Markdown (with descriptions as comments and defaults as values)};
    J --> K{Write combined Markdown to docs/pages/references/configuration.mdx};
    K --> L[End];
```

**Detailed Steps for the Tool:**

1.  **Project Setup**:
    *   Create a new Go module for the tool.
2.  **File & AST Parsing**:
    *   Use the `go/parser` and `go/ast` packages to parse all `.go` files within the `internal/config/` directory.
3.  **Struct Traversal & Information Extraction**:
    *   Identify all structs that represent configuration blocks.
    *   Recursively traverse all fields of these structs.
    *   For each field, extract:
        *   Go field name.
        *   Value of the `yaml` tag.
        *   Value of the `env` tag.
        *   Value of the `desc` tag.
        *   Value of the `required` tag.
        *   The field's data type.
4.  **Default Value Determination**:
    *   Create an instance of the main `config.Config` struct.
    *   Call its `InitDefaults()` method.
    *   Use the `reflect` package to access the actual default value of each field (and sub-field) from this initialized instance.
5.  **Handling "One-of" Types (e.g., `MQsConfig`)**:
    *   The tool needs to recognize patterns where a struct embeds multiple mutually exclusive configuration types (e.g., `AWSSQSConfig`, `GCPPubSubConfig`, `RabbitMQConfig` within `MQsConfig`).
    *   The generated YAML documentation should clearly indicate these are choices (e.g., using comments like `# Choose one of the following MQ providers:`).
6.  **Markdown Generation**:
    *   **Environment Variables Table**:
        *   Iterate through all collected field information.
        *   For fields that have an `env` tag, create a row in the Markdown table: `| ENV_VARIABLE | Description | Default Value | Required |`.
        *   The "Required" column will be:
            *   "Yes" if `required:"Y"`.
            *   "No" if `required:"N"`.
            *   "Conditional" (or "See Description") if `required:"C"`. The `desc` field (which is also part of the table) will contain the detailed condition.
    *   **YAML Configuration Section**:
        *   Reconstruct the YAML structure based on the struct definitions and `yaml` tags.
        *   Use the `desc` tag content as comments for each YAML key.
        *   Use the extracted default values as the example values in the YAML.
        *   Handle nested structs correctly to show proper indentation and structure.
        *   For "one-of" types, add appropriate comments.
7.  **Output**:
    *   Overwrite the content of `docs/pages/references/configuration.mdx` with the newly generated Markdown.

## Phase 3: Integration into Development Workflow

1.  **Ease of Execution**:
    *   Provide a simple way to run the `configdocsgen` tool (e.g., a `go run` command, a Makefile target like `make docs-config`).
2.  **Automation (Recommended)**:
    *   Use a `go:generate` directive in one of the `internal/config/` files (e.g., in `config.go`):
        ```go
        //go:generate go run ../../cmd/configdocsgen/main.go -output ../../docs/pages/references/configuration.mdx
        package config
        ```
    This allows regeneration by running `go generate ./...` from the project root.
3.  **Developer Guidance**:
    *   Ensure developers know to run this tool (or `go generate`) after making any changes to configuration structs or their tags.
    *   Consider adding a pre-commit hook or a CI check to ensure documentation is up-to-date.

## Benefits of this Approach

*   **Single Source of Truth**: The Go code becomes the definitive source for all configuration details.
*   **Consistency**: Documentation will accurately reflect the code.
*   **Reduced Brittleness**: Changes in code (new fields, changed tags, different defaults) are automatically propagated to the documentation.
*   **Developer Efficiency**: Reduces the manual effort and potential for errors in keeping docs updated.