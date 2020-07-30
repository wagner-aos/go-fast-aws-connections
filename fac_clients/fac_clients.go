package facclients

import (
	awscredentials "github.com/wagner-aos/go-fast-aws-connections/fac_credentials"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

// S3 - client
func S3(region, profile, endpoint string) s3iface.S3API {
	//c := awscredentials.Clients{}
	session, config := awscredentials.GetCredentialsWithChain(region, profile)
	if len(endpoint) > 0 {
		config.Endpoint = &endpoint
	}
	return s3.New(session, config)
}

// SQS - client
func SQS(region, profile, endpoint string) sqsiface.SQSAPI {
	session, config := awscredentials.GetCredentialsWithChain(region, profile)
	if len(endpoint) > 0 {
		config.Endpoint = &endpoint
	}
	return sqs.New(session, config)
}

// DynamoDB - client
func DynamoDB(region, profile, endpoint string) dynamodbiface.DynamoDBAPI {
	session, config := awscredentials.GetCredentialsWithChain(region, profile)
	if len(endpoint) > 0 {
		config.Endpoint = &endpoint
	}
	return dynamodb.New(session, config)
}
