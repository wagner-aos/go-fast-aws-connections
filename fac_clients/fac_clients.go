package facclients

import (
	awscredentials "go-fast-aws-connections/fac_credentials"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// S3 - client
func S3(profile string) s3iface.S3API {
	c := awscredentials.Clients{}
	return s3.New(c.SessionWithProfile(profile), c.Config(profile))
}

// SQS - client
func SQS(profile string) sqsiface.SQSAPI {
	c := awscredentials.Clients{}
	return sqs.New(c.SessionWithProfile(profile), c.Config(profile))
}

// DynamoDB - client
func DynamoDB(profile string) dynamodbiface.DynamoDBAPI {
	c := awscredentials.Clients{}
	return dynamodb.New(c.SessionWithProfile(profile), c.Config(profile))
}
