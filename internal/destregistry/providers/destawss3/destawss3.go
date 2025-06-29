package destawss3

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	awssdk "github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	awscreds "github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/hookdeck/outpost/internal/destregistry"
	"github.com/hookdeck/outpost/internal/destregistry/metadata"
	"github.com/hookdeck/outpost/internal/models"
)

// Config for S3 destination
type AWSS3Config struct {
	Bucket string
	Region string
	Prefix string
}

// Credentials for S3
type AWSS3Credentials struct {
	Key     string
	Secret  string
	Session string
}

// Provider implementation
type AWSS3Provider struct {
	*destregistry.BaseProvider
}

var _ destregistry.Provider = (*AWSS3Provider)(nil)

// New creates a new AWSS3Provider
func New(loader metadata.MetadataLoader) (*AWSS3Provider, error) {
	base, err := destregistry.NewBaseProvider(loader, "aws_s3")
	if err != nil {
		return nil, err
	}

	return &AWSS3Provider{BaseProvider: base}, nil
}

func (p *AWSS3Provider) Validate(ctx context.Context, destination *models.Destination) error {
	_, _, err := p.resolveConfig(ctx, destination)
	return err
}

func (p *AWSS3Provider) CreatePublisher(ctx context.Context, destination *models.Destination) (destregistry.Publisher, error) {
	cfg, creds, err := p.resolveConfig(ctx, destination)
	if err != nil {
		return nil, err
	}

	sdkConfig, err := awsconfig.LoadDefaultConfig(ctx,
		awsconfig.WithCredentialsProvider(awscreds.NewStaticCredentialsProvider(
			creds.Key,
			creds.Secret,
			creds.Session,
		)),
		awsconfig.WithRegion(cfg.Region),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(sdkConfig)

	return &AWSS3Publisher{
		BasePublisher: &destregistry.BasePublisher{},
		client:        client,
		bucket:        cfg.Bucket,
		prefix:        cfg.Prefix,
	}, nil
}

func (p *AWSS3Provider) resolveConfig(ctx context.Context, destination *models.Destination) (*AWSS3Config, *AWSS3Credentials, error) {
	if err := p.BaseProvider.Validate(ctx, destination); err != nil {
		return nil, nil, err
	}

	return &AWSS3Config{
			Bucket: destination.Config["bucket"],
			Region: destination.Config["region"],
			Prefix: destination.Config["prefix"],
		}, &AWSS3Credentials{
			Key:     destination.Credentials["key"],
			Secret:  destination.Credentials["secret"],
			Session: destination.Credentials["session"],
		}, nil
}

func (p *AWSS3Provider) ComputeTarget(destination *models.Destination) destregistry.DestinationTarget {
	bucket := destination.Config["bucket"]
	region := destination.Config["region"]
	return destregistry.DestinationTarget{
		Target:    fmt.Sprintf("%s in %s", bucket, region),
		TargetURL: "",
	}
}

// Publisher implementation

type AWSS3Publisher struct {
	*destregistry.BasePublisher
	client *s3.Client
	bucket string
	prefix string
}

func (p *AWSS3Publisher) Close() error {
	p.BasePublisher.StartClose()
	return nil
}

func (p *AWSS3Publisher) Format(ctx context.Context, event *models.Event) (*s3.PutObjectInput, error) {
	payload := map[string]interface{}{
		"metadata": p.BasePublisher.MakeMetadata(event, time.Now()),
		"data":     event.Data,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	key := fmt.Sprintf("%s%s_%s", p.prefix, time.Now().UTC().Format("20060102T150405Z"), event.ID)

	return &s3.PutObjectInput{
		Bucket: awssdk.String(p.bucket),
		Key:    awssdk.String(key),
		Body:   bytes.NewReader(data),
	}, nil
}

func (p *AWSS3Publisher) Publish(ctx context.Context, event *models.Event) (*destregistry.Delivery, error) {
	if err := p.BasePublisher.StartPublish(); err != nil {
		return nil, err
	}
	defer p.BasePublisher.FinishPublish()

	input, err := p.Format(ctx, event)
	if err != nil {
		return nil, destregistry.NewErrDestinationPublishAttempt(
			err,
			"aws_s3",
			map[string]interface{}{"error": "format_failed", "message": err.Error()},
		)
	}

	_, err = p.client.PutObject(ctx, input)
	if err != nil {
		return &destregistry.Delivery{
				Status: "failed",
				Code:   "ERR",
				Response: map[string]interface{}{
					"error": err.Error(),
				},
			}, destregistry.NewErrDestinationPublishAttempt(err, "aws_s3", map[string]interface{}{
				"error": err.Error(),
			})
	}

	return &destregistry.Delivery{
		Status: "success",
		Code:   "OK",
		Response: map[string]interface{}{
			"bucket": p.bucket,
			"key":    *input.Key,
		},
	}, nil
}

// NewAWSS3Publisher exposed for testing
func NewAWSS3Publisher(client *s3.Client, bucket, prefix string) *AWSS3Publisher {
	return &AWSS3Publisher{
		BasePublisher: &destregistry.BasePublisher{},
		client:        client,
		bucket:        bucket,
		prefix:        prefix,
	}
}
