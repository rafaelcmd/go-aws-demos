package s3basics

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

type S3Client struct {
	S3Client *s3.Client
}

func (s3c S3Client) ListBuckets() []types.Bucket {
	result, err := s3c.S3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	var buckets []types.Bucket
	if err != nil {
		log.Printf("Couldn't list buckets for your account. Here's why: %v\n", err)
	} else {
		buckets = result.Buckets
	}
	for _, bucket := range buckets {
		log.Printf("Bucket: %s\n", *bucket.Name)
	}
	return buckets
}

func (s3c S3Client) CreateBucket(bucketName string) {
	currentBucketsList := s3c.ListBuckets()
	for _, bucket := range currentBucketsList {
		if *bucket.Name == bucketName {
			log.Printf("Bucket %s already exists", bucketName)
			return
		}
	}

	result, err := s3c.S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		log.Printf("Couldn't create bucket %s. Here's why: %v\n", bucketName, err)
	} else {
		log.Printf("Bucket created successfully in %s", *result.Location)
	}
}

func (s3c S3Client) DeleteBucket(bucketName string) {
	_, err := s3c.S3Client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		log.Printf("Couldn't delete bucket %s. Here's why: %v\n", bucketName, err)
	} else {
		log.Printf("Bucket %s deleted successfully", bucketName)
	}
}
