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
func S3(region string, profile string) s3iface.S3API {
	//c := awscredentials.Clients{}
	return s3.New(awscredentials.GetCredentialsWithChain(region, profile))
}

// SQS - client
func SQS(region string, profile string) sqsiface.SQSAPI {
	//c := awscredentials.Clients{}
	return sqs.New(awscredentials.GetCredentialsWithChain(region, profile))
}

// DynamoDB - client
func DynamoDB(region string, profile string) dynamodbiface.DynamoDBAPI {
	//c := awscredentials.Clients{}
	return dynamodb.New(awscredentials.GetCredentialsWithChain(region, profile))
}
