package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// Used SDK https://github.com/aws/aws-sdk-go

func createAWSSession(AWSRegion, AWSKeyID, AWSKeySecret, AWSSessionToken string) (*session.Session, error) {
	if (AWSRegion != "") && (AWSKeyID != "") && (AWSKeySecret != "") {
		return session.NewSession(&aws.Config{
			Region:      aws.String(AWSRegion),
			Credentials: credentials.NewStaticCredentials(AWSKeyID, AWSKeySecret, AWSSessionToken),
		})
	} else {
		// Load creds from environment, shared credentials (~/.aws/credentials),
		// or EC2 Instance Role.
		return session.NewSession()
	}
}

func getAWSSecretString(secretName, AWSRegion, AWSKeyID, AWSKeySecret, AWSSessionToken string) (string, error) {
	var secret string
	awsSession, err := createAWSSession(AWSRegion, AWSKeyID, AWSKeySecret, AWSSessionToken)
	if err != nil {
		return secret, err
	}

	awsService := secretsmanager.New(awsSession)
	awsRequest := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	awsResponse, err := awsService.GetSecretValue(awsRequest)
	if err != nil {
		return secret, err
	}

	secret = *awsResponse.SecretString

	return secret, err
}
