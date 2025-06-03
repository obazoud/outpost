package destawssqs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeAWSSQSConsoleURL(t *testing.T) {
	tests := []struct {
		name     string
		queueURL string
		want     string
	}{
		{
			name:     "basic us-east-1",
			queueURL: "https://sqs.us-east-1.amazonaws.com/123456789012/test-queue",
			want:     "https://us-east-1.console.aws.amazon.com/sqs/v3/home?region=us-east-1#/queues/https%3A%2F%2Fsqs.us-east-1.amazonaws.com%2F123456789012%2Ftest-queue",
		},
		{
			name:     "different region eu-west-1",
			queueURL: "https://sqs.eu-west-1.amazonaws.com/987654321098/production-queue",
			want:     "https://eu-west-1.console.aws.amazon.com/sqs/v3/home?region=eu-west-1#/queues/https%3A%2F%2Fsqs.eu-west-1.amazonaws.com%2F987654321098%2Fproduction-queue",
		},
		{
			name:     "queue name with special characters",
			queueURL: "https://sqs.ap-southeast-2.amazonaws.com/111222333444/my-queue-with-dashes",
			want:     "https://ap-southeast-2.console.aws.amazon.com/sqs/v3/home?region=ap-southeast-2#/queues/https%3A%2F%2Fsqs.ap-southeast-2.amazonaws.com%2F111222333444%2Fmy-queue-with-dashes",
		},
		{
			name:     "localstack URL should return empty",
			queueURL: "http://localhost:4566/000000000000/test-queue",
			want:     "",
		},
		{
			name:     "non-SQS AWS URL should return empty",
			queueURL: "https://s3.us-east-1.amazonaws.com/my-bucket/file.txt",
			want:     "",
		},
		{
			name:     "invalid URL should return empty",
			queueURL: "not-a-url",
			want:     "",
		},
		{
			name:     "empty string should return empty",
			queueURL: "",
			want:     "",
		},
		{
			name:     "SQS compatible service should return empty",
			queueURL: "https://sqs.custom-domain.com/123456789012/queue",
			want:     "",
		},
		{
			name:     "SQS URL with only account ID still generates console URL",
			queueURL: "https://sqs.us-east-1.amazonaws.com/123456789012",
			want:     "https://us-east-1.console.aws.amazon.com/sqs/v3/home?region=us-east-1#/queues/https%3A%2F%2Fsqs.us-east-1.amazonaws.com%2F123456789012",
		},
		{
			name:     "SQS URL with only queue name still generates console URL",
			queueURL: "https://sqs.us-east-1.amazonaws.com/test-queue",
			want:     "https://us-east-1.console.aws.amazon.com/sqs/v3/home?region=us-east-1#/queues/https%3A%2F%2Fsqs.us-east-1.amazonaws.com%2Ftest-queue",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := makeAWSSQSConsoleURL(test.queueURL)
			assert.Equal(t, test.want, got)
		})
	}
}
