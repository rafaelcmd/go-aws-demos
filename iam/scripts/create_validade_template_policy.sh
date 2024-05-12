#!/bin/bash

#This script creates the validate template policy

#Check if POLICY_ARN is set in the environment
if [ -z "$POLICY_ARN" ]; then
	echo "POLICY_ARN is not set. Please set it before running this script."
	exit 1
fi

#Create the policy
aws iam create-policy --policy-name ValidateTemplatePermission --policy-document file://policies/validate_template.json

#Attach the policy to the user
aws iam attach-user-policy --user-name rafael --policy-arn "$POLICY_ARN"