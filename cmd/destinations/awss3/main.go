package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	AWSRegion         = "us-east-1"
	AWSEndpoint       = "http://localhost:4566"
	AWSCredentials    = "test:test:"
	DestinationBucket = "destination_s3_bucket"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	ctx := context.Background()

	credsParts := strings.Split(AWSCredentials, ":")
	if len(credsParts) != 3 {
		return fmt.Errorf("invalid AWS credentials format")
	}

	awsCfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(AWSRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			credsParts[0], credsParts[1], credsParts[2],
		)),
	)
	if err != nil {
		return err
	}

	s3Client := s3.NewFromConfig(awsCfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(AWSEndpoint)
		o.UsePathStyle = true
	})

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	seen := make(map[string]struct{})
	if err := checkForNewObjects(ctx, s3Client, seen, false); err != nil {
		return err
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	go func() {
		for {
			select {
			case <-ticker.C:
				if err := checkForNewObjects(ctx, s3Client, seen, true); err != nil {
					log.Printf("[*] error listing bucket: %v", err)
				}
			case <-termChan:
				return
			}
		}
	}()

	log.Printf("[*] Ready to monitor S3 bucket.\n\tEndpoint: %s\n\tBucket: %s", AWSEndpoint, DestinationBucket)
	log.Printf("[*] Waiting for logs. To exit press CTRL+C")
	<-termChan
	return nil
}

func checkForNewObjects(ctx context.Context, client *s3.Client, seen map[string]struct{}, logNew bool) error {
	var token *string
	for {
		out, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
			Bucket:            aws.String(DestinationBucket),
			ContinuationToken: token,
		})
		if err != nil {
			return err
		}
		for _, obj := range out.Contents {
			key := aws.ToString(obj.Key)
			if _, ok := seen[key]; !ok {
				seen[key] = struct{}{}
				if logNew {
					log.Printf("[x] New object: %s (%d bytes)", key, obj.Size)
				}
			}
		}
		if out.IsTruncated == nil || !*out.IsTruncated {
			break
		}
		token = out.NextContinuationToken
	}
	return nil
}
