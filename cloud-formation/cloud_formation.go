package cloudformation

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	types2 "github.com/aws/aws-sdk-go-v2/service/cloudformation/types"
)

type CloudFormation struct {
	CFClient *cloudformation.Client
}

func (cf CloudFormation) CreateBucketWithCloudFormation(templateBody []byte) {
	var result, err = cf.CFClient.CreateStack(context.TODO(), &cloudformation.CreateStackInput{
		StackName:    aws.String("S3BucketTestStack"),
		TemplateBody: aws.String(string(templateBody)),
		Capabilities: []types2.Capability{"CAPABILITY_NAMED_IAM"},
	})
	if err != nil {
		log.Printf("Couldn't create stack. Here's why: %v\n", err)
	}
	log.Printf("Stack %s created successfully", *result.StackId)
}

func (cf CloudFormation) CreateBucketWithStaticWebSite(templateBody []byte) {
	var result, err = cf.CFClient.CreateStack(context.TODO(), &cloudformation.CreateStackInput{
		StackName:    aws.String("S3BucketStaticWebSite"),
		TemplateBody: aws.String(string(templateBody)),
		Capabilities: []types2.Capability{"CAPABILITY_NAMED_IAM"},
	})
	if err != nil {
		log.Printf("Couldn't create stack. Here's why: %v\n", err)
	}
	log.Printf("Stack %s created successfully", *result.StackId)
}
