package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	cf "github.com/rafaelcmd/go-aws-demos/cloud-formation"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("Couldn't load configuration: %v", err)
	}

	cfClient := cloudformation.NewFromConfig(cfg)
	cf := cf.CloudFormation{CFClient: cfClient}

	templateBody, err := os.ReadFile("create_bucket.yaml")
	if err != nil {
		log.Fatalf("Couldn't load template: %v", err)
	}

	cf.CreateBucketWithCloudFormation(templateBody)

	staticWebSiteTemplate, err := os.ReadFile("create_bucket_static_website.yaml")
	if err != nil {
		log.Fatalf("Couldn't load template: %v", err)
	}

	cf.CreateBucketWithStaticWebSite(staticWebSiteTemplate)
}
