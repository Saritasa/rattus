package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// Used SDK https://github.com/aws/aws-sdk-go

func createAWSSession(
	AWSRegion,
	AWSKeyID,
	AWSKeySecret,
	AWSSessionToken string,
	Debug bool,
) (*session.Session, error) {
	var awsCredentials *credentials.Credentials

	if (AWSKeyID != "") && (AWSKeySecret != "") && (AWSSessionToken != "") {
		awsCredentials = credentials.NewStaticCredentials(
			AWSKeyID,
			AWSKeySecret,
			AWSSessionToken,
		)
	} else if (AWSKeyID != "") && (AWSKeySecret != "") {
		awsCredentials = credentials.NewStaticCredentials(
			AWSKeyID,
			AWSKeySecret,
			"",
		)
	}

	var awsRegion *string
	if AWSRegion != "" {
		awsRegion = aws.String(AWSRegion)
	}

	var logLevel aws.LogLevelType = aws.LogOff
	if Debug {
		logLevel = aws.LogDebug
	}
	return session.NewSession(&aws.Config{
		Region:                        awsRegion,
		Credentials:                   awsCredentials,
		CredentialsChainVerboseErrors: aws.Bool(Debug),
		LogLevel:                      aws.LogLevel(logLevel),
	})
}

func getAWSSecretString(
	secretName,
	AWSRegion,
	AWSKeyID,
	AWSKeySecret,
	AWSSessionToken string,
	Debug bool,
) (string, error) {
	var secret string
	awsSession, err := createAWSSession(AWSRegion, AWSKeyID, AWSKeySecret, AWSSessionToken, Debug)
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
