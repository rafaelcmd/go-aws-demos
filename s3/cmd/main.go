package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3basics "github.com/rafaelcmd/go-aws-demos/s3"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("Couldn't load configuration: %v", err)
	}

	s3Client := s3.NewFromConfig(cfg)
	s3Basics := s3basics.S3Client{S3Client: s3Client}

	bucketName := "bucket-test-aws-sdk-go-rafaelcmd"

	s3Basics.ListBuckets()
	s3Basics.CreateBucket(bucketName)
	s3Basics.DeleteBucket(bucketName)
}
