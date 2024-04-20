package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	"log"
	"os"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("Couldn't load configuration: %v", err)
	}

	//s3Client := s3.NewFromConfig(cfg)
	//s3Basics := BucketBasics{S3Client: s3Client}

	//bucketName := "bucket-test-aws-sdk-go-rafaelcmd"

	//s3Basics.ListBuckets()
	//s3Basics.CreateBucket(bucketName)
	//s3Basics.DeleteBucket(bucketName)

	cfClient := cloudformation.NewFromConfig(cfg)
	cfBasics := BucketBasics{CFClient: cfClient}

	//templateBody, err := os.ReadFile("create_bucket.yaml")
	staticWebSiteTemplate, err := os.ReadFile("create_bucket_static_website.yaml")

	//cfBasics.CreateBucketWithCloudFormation(staticWebSiteTemplate)
	cfBasics.CreateBucketWithStaticWebSite(staticWebSiteTemplate)
}
