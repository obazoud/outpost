package config

import (
	"errors"
	"fmt"

	"github.com/hookdeck/outpost/internal/otel"
	v "github.com/spf13/viper"
)

type OpenTelemetryTypeConfig struct {
	Exporter string `yaml:"exporter" env:"OTEL_EXPORTER" desc:"Specifies the OTLP exporter to use for this telemetry type (e.g., 'otlp'). Typically used with environment variables like OTEL_EXPORTER_OTLP_TRACES_ENDPOINT." required:"C"`
	Protocol string `yaml:"protocol" env:"OTEL_PROTOCOL" desc:"Specifies the OTLP protocol ('grpc' or 'http') for this telemetry type. Typically used with environment variables like OTEL_EXPORTER_OTLP_TRACES_PROTOCOL." required:"C"`
}

type OpenTelemetryConfig struct {
	ServiceName string                  `yaml:"service_name" env:"OTEL_SERVICE_NAME" desc:"The service name reported to OpenTelemetry. If set, OpenTelemetry will be enabled." required:"N"`
	Traces      OpenTelemetryTypeConfig `yaml:"traces" desc:"OpenTelemetry configuration specific to traces."`
	Metrics     OpenTelemetryTypeConfig `yaml:"metrics" desc:"OpenTelemetry configuration specific to metrics."`
	Logs        OpenTelemetryTypeConfig `yaml:"logs" desc:"OpenTelemetry configuration specific to logs."`
}

const (
	OTelProtocolGRPC = "grpc"
	OTelProtocolHTTP = "http"
)

var validOTelProtocols = map[string]bool{
	OTelProtocolGRPC: true,
	OTelProtocolHTTP: true,
}

var ErrInvalidOTelProtocol = errors.New("config validation error: invalid OpenTelemetry protocol, must be 'grpc' or 'http'")

func validateOTelProtocol(protocol string) error {
	if protocol == "" {
		return nil // Empty protocol will use default
	}
	if !validOTelProtocols[protocol] {
		return ErrInvalidOTelProtocol
	}
	return nil
}

func getProtocol(viper *v.Viper, telemetryType string) string {
	// Check type-specific protocol first
	protocol := viper.GetString(fmt.Sprintf("OTEL_EXPORTER_OTLP_%s_PROTOCOL", telemetryType))
	if protocol == "" {
		// Fall back to generic protocol
		protocol = viper.GetString("OTEL_EXPORTER_OTLP_PROTOCOL")
	}
	if protocol == "" {
		// Default to gRPC if not specified
		protocol = "grpc"
	}
	return protocol
}

func (c *OpenTelemetryConfig) Validate() error {
	if c.ServiceName == "" {
		return nil // OpenTelemetry is optional
	}

	if err := validateOTelProtocol(c.Traces.Protocol); err != nil {
		return err
	}
	if err := validateOTelProtocol(c.Metrics.Protocol); err != nil {
		return err
	}
	if err := validateOTelProtocol(c.Logs.Protocol); err != nil {
		return err
	}

	return nil
}

func (c *OpenTelemetryConfig) ToConfig() *otel.OpenTelemetryConfig {
	if c.ServiceName == "" {
		return nil
	}

	// Set default protocol if not specified
	getProtocolWithDefault := func(p string) string {
		if p == "" {
			return OTelProtocolGRPC
		}
		return p
	}

	return &otel.OpenTelemetryConfig{
		Traces: &otel.OpenTelemetryTypeConfig{
			Exporter: c.Traces.Exporter,
			Protocol: getProtocolWithDefault(c.Traces.Protocol),
		},
		Metrics: &otel.OpenTelemetryTypeConfig{
			Exporter: c.Metrics.Exporter,
			Protocol: getProtocolWithDefault(c.Metrics.Protocol),
		},
		Logs: &otel.OpenTelemetryTypeConfig{
			Exporter: c.Logs.Exporter,
			Protocol: getProtocolWithDefault(c.Logs.Protocol),
		},
	}
}

func (c *OpenTelemetryConfig) GetServiceName() string {
	if c == nil {
		return ""
	}
	return c.ServiceName
}
