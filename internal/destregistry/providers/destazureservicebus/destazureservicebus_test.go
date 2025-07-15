package destazureservicebus_test

import (
	"testing"

	"github.com/hookdeck/outpost/internal/destregistry/providers/destazureservicebus"
	"github.com/hookdeck/outpost/internal/models"
	"github.com/hookdeck/outpost/internal/util/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComputeTarget(t *testing.T) {
	provider, err := destazureservicebus.New(testutil.Registry.MetadataLoader())
	require.NoError(t, err)

	tests := []struct {
		name           string
		destination    models.Destination
		expectedTarget string
	}{
		{
			name: "with valid connection string and name",
			destination: models.Destination{
				Config: map[string]string{
					"name": "my-queue",
				},
				Credentials: map[string]string{
					"connection_string": "Endpoint=sb://mynamespace.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=abcd1234",
				},
			},
			expectedTarget: "mynamespace/my-queue",
		},
		{
			name: "with different namespace format",
			destination: models.Destination{
				Config: map[string]string{
					"name": "my-topic",
				},
				Credentials: map[string]string{
					"connection_string": "Endpoint=sb://test-namespace.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=xyz789",
				},
			},
			expectedTarget: "test-namespace/my-topic",
		},
		{
			name: "with missing name config",
			destination: models.Destination{
				Config: map[string]string{},
				Credentials: map[string]string{
					"connection_string": "Endpoint=sb://mynamespace.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=abcd1234",
				},
			},
			expectedTarget: "",
		},
		{
			name: "with invalid connection string format",
			destination: models.Destination{
				Config: map[string]string{
					"name": "my-queue",
				},
				Credentials: map[string]string{
					"connection_string": "invalid-connection-string",
				},
			},
			expectedTarget: "my-queue", // Falls back to just the name
		},
		{
			name: "with missing connection string",
			destination: models.Destination{
				Config: map[string]string{
					"name": "my-queue",
				},
				Credentials: map[string]string{},
			},
			expectedTarget: "my-queue", // Falls back to just the name
		},
		{
			name: "with connection string missing Endpoint",
			destination: models.Destination{
				Config: map[string]string{
					"name": "my-queue",
				},
				Credentials: map[string]string{
					"connection_string": "SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=abcd1234",
				},
			},
			expectedTarget: "my-queue", // Falls back to just the name
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := provider.ComputeTarget(&tt.destination)
			assert.Equal(t, tt.expectedTarget, result.Target)
			assert.Empty(t, result.TargetURL) // TargetURL should always be empty for now
		})
	}
}

func TestParseNamespaceFromConnectionString(t *testing.T) {
	tests := []struct {
		name              string
		connectionString  string
		expectedNamespace string
	}{
		{
			name:              "standard connection string",
			connectionString:  "Endpoint=sb://mynamespace.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=abcd1234",
			expectedNamespace: "mynamespace",
		},
		{
			name:              "connection string with hyphenated namespace",
			connectionString:  "Endpoint=sb://my-test-namespace.servicebus.windows.net/;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=xyz789",
			expectedNamespace: "my-test-namespace",
		},
		{
			name:              "connection string with different order",
			connectionString:  "SharedAccessKeyName=RootManageSharedAccessKey;Endpoint=sb://namespace123.servicebus.windows.net/;SharedAccessKey=key123",
			expectedNamespace: "namespace123",
		},
		{
			name:              "invalid connection string",
			connectionString:  "invalid-string",
			expectedNamespace: "",
		},
		{
			name:              "missing endpoint",
			connectionString:  "SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=abcd1234",
			expectedNamespace: "",
		},
		{
			name:              "malformed endpoint",
			connectionString:  "Endpoint=invalid-endpoint;SharedAccessKeyName=RootManageSharedAccessKey;SharedAccessKey=abcd1234",
			expectedNamespace: "",
		},
		{
			name:              "empty connection string",
			connectionString:  "",
			expectedNamespace: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// We need to test the parseNamespaceFromConnectionString function
			// Since it's not exported, we test it indirectly through ComputeTarget
			provider, err := destazureservicebus.New(testutil.Registry.MetadataLoader())
			require.NoError(t, err)

			dest := models.Destination{
				Config: map[string]string{
					"name": "test-entity",
				},
				Credentials: map[string]string{
					"connection_string": tt.connectionString,
				},
			}

			result := provider.ComputeTarget(&dest)
			if tt.expectedNamespace == "" {
				assert.Equal(t, "test-entity", result.Target)
			} else {
				assert.Equal(t, tt.expectedNamespace+"/test-entity", result.Target)
			}
		})
	}
}
