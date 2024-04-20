#!/bin/bash

#Check if POLICY_ARN is set in the environment
if [ -z "$CF_POLICY_ARN" ]; then
	echo "CF_POLICY_ARN is not set. Please set it before running this script."
	exit 1
fi

#Create the IAM Role
aws iam create-role --role-name CloudFormationS3AccessRole --assume-role-policy-document file://policies/cloud_formation_s3_trust_relationship_policy.json

#Create the S3 Access Policy
aws iam create-policy --policy-name CloudFormationS3Policy --policy-document file://policies/s3_policy.json

#Attach the S3 Access Policy to the IAM Role
aws iam attach-role-policy --role-name CloudFormationS3AccessRole --policy-arn "$CF_POLICY_ARN"