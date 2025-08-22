//go:generate go run ../../cmd/configdocsgen/main.go -input-dir . -output-file ../../docs/pages/references/configuration.mdx
package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/caarlos0/env/v9"
	"github.com/hookdeck/outpost/internal/migrator"
	"github.com/hookdeck/outpost/internal/redis"
	"github.com/hookdeck/outpost/internal/telemetry"
	"github.com/hookdeck/outpost/internal/version"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

const (
	Namespace = "Outpost"
)

func getConfigLocations() []string {
	return []string{
		// Relative paths
		".env",
		".outpost.yaml",
		"config/outpost.yaml",
		"config/outpost/config.yaml",
		"config/outpost/.env",

		// Container-friendly absolute paths
		"/config/outpost.yaml",
		"/config/outpost/config.yaml",
		"/config/outpost/.env",
	}
}

type Config struct {
	validated  bool   // tracks whether Validate() has been called successfully
	configPath string // stores the path of the config file used

	Service       string              `yaml:"service" env:"SERVICE" desc:"Specifies the service type to run. Valid values: 'api', 'log', 'delivery', or empty/all for singular mode (runs all services)." required:"N"`
	LogLevel      string              `yaml:"log_level" env:"LOG_LEVEL" desc:"Defines the verbosity of application logs. Common values: 'trace', 'debug', 'info', 'warn', 'error'." required:"N"`
	AuditLog      bool                `yaml:"audit_log" env:"AUDIT_LOG" desc:"Enables or disables audit logging for significant events." required:"N"`
	OpenTelemetry OpenTelemetryConfig `yaml:"otel"`
	Telemetry     TelemetryConfig     `yaml:"telemetry"`

	// API
	APIPort      int    `yaml:"api_port" env:"API_PORT" desc:"Port number for the API server to listen on." required:"N"`
	APIKey       string `yaml:"api_key" env:"API_KEY" desc:"API key for authenticating requests to the Outpost API." required:"Y"`
	APIJWTSecret string `yaml:"api_jwt_secret" env:"API_JWT_SECRET" desc:"Secret key for signing and verifying JWTs if JWT authentication is used for the API." required:"Y"`
	GinMode      string `yaml:"gin_mode" env:"GIN_MODE" desc:"Sets the Gin framework mode (e.g., 'debug', 'release', 'test'). See Gin documentation for details." required:"N"`

	// Application
	AESEncryptionSecret string   `yaml:"aes_encryption_secret" env:"AES_ENCRYPTION_SECRET" desc:"A 16, 24, or 32 byte secret key used for AES encryption of sensitive data at rest." required:"Y"`
	Topics              []string `yaml:"topics" env:"TOPICS" envSeparator:"," desc:"Comma-separated list of topics that this Outpost instance should subscribe to for event processing." required:"N"`
	OrganizationName    string   `yaml:"organization_name" env:"ORGANIZATION_NAME" desc:"Name of the organization, used for display purposes and potentially in user agent strings." required:"N"`
	HTTPUserAgent       string   `yaml:"http_user_agent" env:"HTTP_USER_AGENT" desc:"Custom HTTP User-Agent string for outgoing webhook deliveries. If unset, a default (OrganizationName/Version) is used." required:"N"`

	// Infrastructure
	Redis RedisConfig `yaml:"redis"`
	// ClickHouse  ClickHouseConfig `yaml:"clickhouse"`
	PostgresURL string     `yaml:"postgres" env:"POSTGRES_URL" desc:"Connection URL for PostgreSQL, used for log storage. Example: 'postgres://user:pass@host:port/dbname?sslmode=disable'." required:"Y"`
	MQs         *MQsConfig `yaml:"mqs"`

	// PublishMQ
	PublishMQ PublishMQConfig `yaml:"publishmq"`

	// Consumers
	PublishMaxConcurrency  int `yaml:"publish_max_concurrency" env:"PUBLISH_MAX_CONCURRENCY" desc:"Maximum number of messages to process concurrently from the publish queue." required:"N"`
	DeliveryMaxConcurrency int `yaml:"delivery_max_concurrency" env:"DELIVERY_MAX_CONCURRENCY" desc:"Maximum number of delivery attempts to process concurrently." required:"N"`
	LogMaxConcurrency      int `yaml:"log_max_concurrency" env:"LOG_MAX_CONCURRENCY" desc:"Maximum number of log writing operations to process concurrently." required:"N"`

	// Delivery Retry
	RetryIntervalSeconds int `yaml:"retry_interval_seconds" env:"RETRY_INTERVAL_SECONDS" desc:"Interval in seconds between delivery retry attempts for failed webhooks." required:"N"`
	RetryMaxLimit        int `yaml:"retry_max_limit" env:"MAX_RETRY_LIMIT" desc:"Maximum number of retry attempts for a single event delivery before giving up." required:"N"`

	// Event Delivery
	MaxDestinationsPerTenant int `yaml:"max_destinations_per_tenant" env:"MAX_DESTINATIONS_PER_TENANT" desc:"Maximum number of destinations allowed per tenant/organization." required:"N"`
	DeliveryTimeoutSeconds   int `yaml:"delivery_timeout_seconds" env:"DELIVERY_TIMEOUT_SECONDS" desc:"Timeout in seconds for HTTP requests made during event delivery to webhook destinations." required:"N"`

	// Destination Registry
	DestinationMetadataPath string `yaml:"destination_metadata_path" env:"DESTINATION_METADATA_PATH" desc:"Path to the directory containing custom destination type definitions. Overrides 'destinations.metadata_path' if set." required:"N"`

	// Log batcher configuration
	LogBatchThresholdSeconds int `yaml:"log_batch_threshold_seconds" env:"LOG_BATCH_THRESHOLD_SECONDS" desc:"Maximum time in seconds to buffer logs before flushing them to storage, if batch size is not reached." required:"N"`
	LogBatchSize             int `yaml:"log_batch_size" env:"LOG_BATCH_SIZE" desc:"Maximum number of log entries to batch together before writing to storage." required:"N"`

	DisableTelemetry bool `yaml:"disable_telemetry" env:"DISABLE_TELEMETRY" desc:"Global flag to disable all telemetry (anonymous usage statistics to Hookdeck and error reporting to Sentry). If true, overrides 'telemetry.disabled'." required:"N"`

	// Destinations
	Destinations DestinationsConfig `yaml:"destinations"`

	// Portal
	Portal PortalConfig `yaml:"portal"`

	// Alert
	Alert AlertConfig `yaml:"alert"`
}

var (
	ErrMismatchedServiceType = errors.New("config validation error: service type mismatch")
	ErrInvalidServiceType    = errors.New("config validation error: invalid service type")
	ErrMissingRedis          = errors.New("config validation error: redis configuration is required")
	ErrMissingLogStorage     = errors.New("config validation error: log storage must be provided")
	ErrMissingMQs            = errors.New("config validation error: message queue configuration is required")
	ErrMissingAESSecret      = errors.New("config validation error: AES encryption secret is required")
	ErrInvalidPortalProxyURL = errors.New("config validation error: invalid portal proxy url")
)

func (c *Config) InitDefaults() {
	c.APIPort = 3333
	c.LogLevel = "info"
	c.AuditLog = true
	c.OpenTelemetry = OpenTelemetryConfig{}
	c.GinMode = "release"
	c.Redis = RedisConfig{
		Host: "127.0.0.1",
		Port: 6379,
	}
	// c.ClickHouse = ClickHouseConfig{
	// 	Database: "outpost",
	// }
	c.MQs = &MQsConfig{
		RabbitMQ: RabbitMQConfig{
			Exchange:      "outpost",
			DeliveryQueue: "outpost-delivery",
			LogQueue:      "outpost-log",
		},
		AWSSQS: AWSSQSConfig{
			DeliveryQueue: "outpost-delivery",
			LogQueue:      "outpost-log",
		},
		GCPPubSub: GCPPubSubConfig{
			DeliveryTopic:        "outpost-delivery",
			DeliverySubscription: "outpost-delivery-sub",
			LogTopic:             "outpost-log",
			LogSubscription:      "outpost-log-sub",
		},
		AzureServiceBus: AzureServiceBusConfig{
			DeliveryTopic:        "outpost-delivery",
			DeliverySubscription: "outpost-delivery-sub",
			LogTopic:             "outpost-log",
			LogSubscription:      "outpost-log-sub",
		},
	}
	c.PublishMaxConcurrency = 1
	c.DeliveryMaxConcurrency = 1
	c.LogMaxConcurrency = 1
	c.RetryIntervalSeconds = 30
	c.RetryMaxLimit = 10
	c.MaxDestinationsPerTenant = 20
	c.DeliveryTimeoutSeconds = 5
	c.LogBatchThresholdSeconds = 10
	c.LogBatchSize = 1000

	// Set defaults for Destinations config
	c.Destinations = DestinationsConfig{
		MetadataPath: "config/outpost/destinations",
		Webhook: DestinationWebhookConfig{
			HeaderPrefix:             "x-outpost-",
			SignatureContentTemplate: "{{.Timestamp.Unix}}.{{.Body}}",
			SignatureHeaderTemplate:  "t={{.Timestamp.Unix}},v0={{.Signatures | join \",\"}}",
			SignatureEncoding:        "hex",
			SignatureAlgorithm:       "hmac-sha256",
		},
		AWSKinesis: DestinationAWSKinesisConfig{
			MetadataInPayload: true,
		},
	}

	c.Alert = AlertConfig{
		CallbackURL:             "",
		ConsecutiveFailureCount: 20,
		AutoDisableDestination:  true,
	}

	c.Telemetry = TelemetryConfig{
		Disabled:          false,
		BatchSize:         100,
		BatchInterval:     5,
		HookdeckSourceURL: "https://hkdk.events/yhk665ljz3rn6l",
		SentryDSN:         "https://examplePublicKey@o0.ingest.sentry.io/0",
	}
}

func (c *Config) parseConfigFile(flagPath string, osInterface OSInterface) error {
	// Get config file path from flag or env
	configPath := flagPath
	if envPath := osInterface.Getenv("CONFIG"); envPath != "" {
		if configPath != "" && configPath != envPath {
			return fmt.Errorf("conflicting config paths: flag=%s env=%s", configPath, envPath)
		}
		configPath = envPath
	}

	// If no explicit config path, try default locations
	if configPath == "" {
		for _, loc := range getConfigLocations() {
			if _, err := osInterface.Stat(loc); err == nil {
				configPath = loc
				break
			}
		}
	}

	if configPath == "" {
		return nil
	}

	data, err := osInterface.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	// Skip empty files
	if len(data) == 0 {
		return nil
	}

	// Store the config path
	c.configPath = configPath

	// Parse based on file extension
	if strings.HasSuffix(strings.ToLower(configPath), ".env") {
		envMap, err := godotenv.Read(configPath)
		if err != nil {
			return fmt.Errorf("error loading .env file: %w", err)
		}
		if err := env.ParseWithOptions(c, env.Options{
			Environment: envMap,
		}); err != nil {
			return fmt.Errorf("error parsing .env file: %w", err)
		}
	} else {
		if err := yaml.Unmarshal(data, c); err != nil {
			return fmt.Errorf("error parsing yaml config: %w", err)
		}
	}
	return nil
}

func (c *Config) parseEnvVariables(osInterface OSInterface) error {
	// For testing, use the mock environment
	if _, ok := osInterface.(*defaultOSImpl); !ok {
		// Build environment map from all env vars
		envMap := make(map[string]string)
		for _, env := range osInterface.Environ() {
			if i := strings.Index(env, "="); i >= 0 {
				envMap[env[:i]] = env[i+1:]
			}
		}
		return env.ParseWithOptions(c, env.Options{Environment: envMap})
	}

	// For real OS, use env.Parse directly
	return env.Parse(c)
}

// GetService returns ServiceType with error checking
func (c *Config) GetService() (ServiceType, error) {
	return ServiceTypeFromString(c.Service)
}

// MustGetService returns ServiceType without error checking - panics if called before validation
func (c *Config) MustGetService() ServiceType {
	if !c.validated {
		panic("MustGetService called before validation")
	}
	// We can skip error checking since validation ensures this is valid
	svc, _ := ServiceTypeFromString(c.Service)
	return svc
}

// ParseWithoutValidation parses the config without validation
func ParseWithoutValidation(flags Flags, osInterface OSInterface) (*Config, error) {
	var config Config

	// Initialize defaults
	config.InitDefaults()

	// Parse config file (lower priority)
	if err := config.parseConfigFile(flags.Config, osInterface); err != nil {
		return nil, err
	}

	// Parse environment variables (highest priority)
	if err := config.parseEnvVariables(osInterface); err != nil {
		return nil, err
	}

	return &config, nil
}

// Parse is the main entry point for parsing and validating config
func Parse(flags Flags) (*Config, error) {
	return ParseWithOS(flags, defaultOS)
}

// ParseWithOS parses and validates config with a custom OS interface
func ParseWithOS(flags Flags, osInterface OSInterface) (*Config, error) {
	config, err := ParseWithoutValidation(flags, osInterface)
	if err != nil {
		return nil, err
	}

	if err := config.Validate(flags); err != nil {
		return nil, err
	}

	return config, nil
}

type RedisConfig struct {
	Host           string `yaml:"host" env:"REDIS_HOST" desc:"Hostname or IP address of the Redis server." required:"Y"`
	Port           int    `yaml:"port" env:"REDIS_PORT" desc:"Port number for the Redis server." required:"Y"`
	Password       string `yaml:"password" env:"REDIS_PASSWORD" desc:"Password for Redis authentication, if required by the server." required:"Y"`
	Database       int    `yaml:"database" env:"REDIS_DATABASE" desc:"Redis database number to select after connecting (ignored in cluster mode)." required:"Y"`
	TLSEnabled     bool   `yaml:"tls_enabled" env:"REDIS_TLS_ENABLED" desc:"Enable TLS encryption for Redis connection (required for Azure Managed Redis)." required:"N"`
	ClusterEnabled bool   `yaml:"cluster_enabled" env:"REDIS_CLUSTER_ENABLED" desc:"Enable Redis cluster mode (required for Azure Managed Redis, optional for self-hosted Redis clusters)." required:"N"`
}

func (c *RedisConfig) ToConfig() *redis.RedisConfig {
	return &redis.RedisConfig{
		Host:           c.Host,
		Port:           c.Port,
		Password:       c.Password,
		Database:       c.Database,
		TLSEnabled:     c.TLSEnabled,
		ClusterEnabled: c.ClusterEnabled,
	}
}

// type ClickHouseConfig struct {
// 	Addr     string `yaml:"addr" env:"CLICKHOUSE_ADDR" desc:"Address (host:port) of the ClickHouse server. Example: 'localhost:9000'. Required if ClickHouse is used for log storage." required:"C"`
// 	Username string `yaml:"username" env:"CLICKHOUSE_USERNAME" desc:"Username for ClickHouse authentication." required:"N"`
// 	Password string `yaml:"password" env:"CLICKHOUSE_PASSWORD" desc:"Password for ClickHouse authentication." required:"N"`
// 	Database string `yaml:"database" env:"CLICKHOUSE_DATABASE" desc:"Database name in ClickHouse to use." required:"N"`
// }

// func (c *ClickHouseConfig) ToConfig() *clickhouse.ClickHouseConfig {
// 	if c.Addr == "" {
// 		return nil
// 	}
// 	return &clickhouse.ClickHouseConfig{
// 		Addr:     c.Addr,
// 		Username: c.Username,
// 		Password: c.Password,
// 		Database: c.Database,
// 	}
// }

type AlertConfig struct {
	CallbackURL             string `yaml:"callback_url" env:"ALERT_CALLBACK_URL" desc:"URL to which Outpost will send a POST request when an alert is triggered (e.g., for destination failures)." required:"N"`
	ConsecutiveFailureCount int    `yaml:"consecutive_failure_count" env:"ALERT_CONSECUTIVE_FAILURE_COUNT" desc:"Number of consecutive delivery failures for a destination before triggering an alert and potentially disabling it." required:"N"`
	AutoDisableDestination  bool   `yaml:"auto_disable_destination" env:"ALERT_AUTO_DISABLE_DESTINATION" desc:"If true, automatically disables a destination after 'consecutive_failure_count' is reached." required:"N"`
}

// ConfigFilePath returns the path of the config file that was used
func (c *Config) ConfigFilePath() string {
	return c.configPath
}

type TelemetryConfig struct {
	Disabled          bool   `yaml:"disabled" env:"DISABLE_TELEMETRY" desc:"Disables telemetry within the 'telemetry' block (Hookdeck usage stats and Sentry). Can be overridden by the global 'disable_telemetry' flag at the root of the configuration." required:"N"`
	BatchSize         int    `yaml:"batch_size" env:"TELEMETRY_BATCH_SIZE" desc:"Maximum number of telemetry events to batch before sending." required:"N"`
	BatchInterval     int    `yaml:"batch_interval" env:"TELEMETRY_BATCH_INTERVAL" desc:"Maximum time in seconds to wait before sending a batch of telemetry events if batch size is not reached." required:"N"`
	HookdeckSourceURL string `yaml:"hookdeck_source_url" env:"TELEMETRY_HOOKDECK_SOURCE_URL" desc:"The Hookdeck Source URL to send anonymous usage telemetry data to. Set to empty to disable sending to Hookdeck." required:"N"`
	SentryDSN         string `yaml:"sentry_dsn" env:"TELEMETRY_SENTRY_DSN" desc:"Sentry DSN for error reporting. If provided and telemetry is not disabled, Sentry integration will be enabled." required:"N"`
}

func (c *TelemetryConfig) ToTelemetryConfig() telemetry.TelemetryConfig {
	return telemetry.TelemetryConfig{
		Disabled:          c.Disabled,
		BatchSize:         c.BatchSize,
		BatchInterval:     c.BatchInterval,
		HookdeckSourceURL: c.HookdeckSourceURL,
		SentryDSN:         c.SentryDSN,
	}
}

func (c *Config) ToTelemetryApplicationInfo() telemetry.ApplicationInfo {
	portalEnabled := c.APIKey != "" && c.APIJWTSecret != ""

	entityStore := "redis"
	logStore := "TODO"
	// if c.ClickHouse.Addr != "" {
	// 	logStore = "clickhouse"
	// }
	if c.PostgresURL != "" {
		logStore = "postgres"
	}

	return telemetry.ApplicationInfo{
		Version:       version.Version(),
		MQ:            c.MQs.GetInfraType(),
		PortalEnabled: portalEnabled,
		EntityStore:   entityStore,
		LogStore:      logStore,
	}
}

// ===== Misc =====

func (c *Config) ToMigratorOpts() migrator.MigrationOpts {
	return migrator.MigrationOpts{
		PG: migrator.MigrationOptsPG{
			URL: c.PostgresURL,
		},
		// CH: migrator.MigrationOptsCH{
		// 	Addr:     c.ClickHouse.Addr,
		// 	Username: c.ClickHouse.Username,
		// 	Password: c.ClickHouse.Password,
		// 	Database: c.ClickHouse.Database,
		// },
	}
}
