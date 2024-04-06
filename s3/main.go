package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
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

	buckets, err := basics.ListBuckets()
	if err != nil {
		log.Fatalf("Couldn't list buckets for your account. Here's why: %v", err)
	}

	fmt.Println("Buckets:")
	for _, bucket := range buckets {
		fmt.Printf("* %s\n", aws.ToString(bucket.Name))
	}
}
