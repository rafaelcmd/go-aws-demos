package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"log"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("Couldn't load configuration: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)

	basics := BucketBasics{S3Client: s3Client}

	bucketName := "bucket-test-aws-sdk-go-rafaelcmd"

	basics.ListBuckets()
	basics.CreateBucket(bucketName)
	basics.DeleteBucket(bucketName)
}
