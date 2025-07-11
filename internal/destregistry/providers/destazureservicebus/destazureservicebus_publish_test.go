package destazureservicebus_test

import (
	"testing"

	"github.com/hookdeck/outpost/internal/destregistry/providers/destazureservicebus"
	testsuite "github.com/hookdeck/outpost/internal/destregistry/testing"
	"github.com/hookdeck/outpost/internal/models"
	"github.com/hookdeck/outpost/internal/util/testinfra"
	"github.com/hookdeck/outpost/internal/util/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// AzureServiceBusConsumer implements testsuite.MessageConsumer
type AzureServiceBusConsumer struct {
	msgChan chan testsuite.Message
	done    chan struct{}
}

func NewAzureServiceBusConsumer() *AzureServiceBusConsumer {
	c := &AzureServiceBusConsumer{
		msgChan: make(chan testsuite.Message),
		done:    make(chan struct{}),
	}
	// Phase 1: Mock consumer that doesn't actually consume
	return c
}

func (c *AzureServiceBusConsumer) Consume() <-chan testsuite.Message {
	return c.msgChan
}

func (c *AzureServiceBusConsumer) Close() error {
	close(c.done)
	return nil
}

type AzureServiceBusAsserter struct{}

func (a *AzureServiceBusAsserter) AssertMessage(t testsuite.TestingT, msg testsuite.Message, event models.Event) {
	// Phase 1: Basic assertion structure
	metadata := msg.Metadata

	// Verify system metadata
	assert.NotEmpty(t, metadata["timestamp"], "timestamp should be present")
	assert.Equal(t, event.ID, metadata["event-id"], "event-id should match")
	assert.Equal(t, event.Topic, metadata["topic"], "topic should match")

	// Verify custom metadata
	for k, v := range event.Metadata {
		assert.Equal(t, v, metadata[k], "metadata key %s should match expected value", k)
	}
}

type AzureServiceBusSuite struct {
	testsuite.PublisherSuite
	consumer *AzureServiceBusConsumer
}

func TestDestinationAzureServiceBusSuite(t *testing.T) {
	suite.Run(t, new(AzureServiceBusSuite))
}

func (s *AzureServiceBusSuite) SetupSuite() {
	t := s.T()
	t.Cleanup(testinfra.Start(t))
	mqConfig := testinfra.GetMQAzureConfig(t, "TestDestinationAzureServiceBusSuite")

	// Create consumer
	s.consumer = NewAzureServiceBusConsumer()

	// Create provider
	provider, err := destazureservicebus.New(testutil.Registry.MetadataLoader())
	require.NoError(t, err)

	// Create destination
	destination := testutil.DestinationFactory.Any(
		testutil.DestinationFactory.WithType("azure_servicebus"),
		testutil.DestinationFactory.WithConfig(map[string]string{
			"topic": mqConfig.AzureServiceBus.Topic,
		}),
		testutil.DestinationFactory.WithCredentials(map[string]string{
			"connection_string": mqConfig.AzureServiceBus.ConnectionString,
		}),
	)

	// Initialize publisher suite with custom asserter
	cfg := testsuite.Config{
		Provider: provider,
		Dest:     &destination,
		Consumer: s.consumer,
		Asserter: &AzureServiceBusAsserter{},
	}
	s.InitSuite(cfg)
}

func (s *AzureServiceBusSuite) TearDownSuite() {
	if s.consumer != nil {
		_ = s.consumer.Close()
	}
}
