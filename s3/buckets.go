package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	types2 "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"log"
)

type BucketBasics struct {
	S3Client *s3.Client
	CFClient *cloudformation.Client
}

func (basics BucketBasics) ListBuckets() []types.Bucket {
	result, err := basics.S3Client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
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

func (basics BucketBasics) CreateBucket(bucketName string) {
	currentBucketsList := basics.ListBuckets()
	for _, bucket := range currentBucketsList {
		if *bucket.Name == bucketName {
			log.Printf("Bucket %s already exists", bucketName)
			return
		}
	}

	result, err := basics.S3Client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		log.Printf("Couldn't create bucket %s. Here's why: %v\n", bucketName, err)
	} else {
		log.Printf("Bucket created successfully in %s", *result.Location)
	}
}

func (basics BucketBasics) DeleteBucket(bucketName string) {
	_, err := basics.S3Client.DeleteBucket(context.TODO(), &s3.DeleteBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		log.Printf("Couldn't delete bucket %s. Here's why: %v\n", bucketName, err)
	} else {
		log.Printf("Bucket %s deleted successfully", bucketName)
	}
}

func (basics BucketBasics) CreateBucketWithCloudFormation(templateBody []byte) {
	var result, err = basics.CFClient.CreateStack(context.TODO(), &cloudformation.CreateStackInput{
		StackName:    aws.String("S3BucketTestStack"),
		TemplateBody: aws.String(string(templateBody)),
		Capabilities: []types2.Capability{"CAPABILITY_NAMED_IAM"},
	})
	if err != nil {
		log.Printf("Couldn't create stack. Here's why: %v\n", err)
	}
	log.Printf("Stack %s created successfully", *result.StackId)
}

func (basics BucketBasics) CreateBucketWithStaticWebSite(templateBody []byte) {
	var result, err = basics.CFClient.CreateStack(context.TODO(), &cloudformation.CreateStackInput{
		StackName:    aws.String("S3BucketStaticWebSite"),
		TemplateBody: aws.String(string(templateBody)),
		Capabilities: []types2.Capability{"CAPABILITY_NAMED_IAM"},
	})
	if err != nil {
		log.Printf("Couldn't create stack. Here's why: %v\n", err)
	}
	log.Printf("Stack %s created successfully", *result.StackId)
}
