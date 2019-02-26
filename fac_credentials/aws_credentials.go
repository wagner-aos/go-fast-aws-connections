package awscredentials

//package main

//https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
//https://github.com/aws/aws-sdk-go
//https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sessions.html

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/aws/aws-sdk-go/service/sts"
)

var awsProfile AWSProfile

//Clients - It store config of each session
type Clients struct {
	session  *session.Session
	configs  map[string]*aws.Config
	profiles map[string]*string
}

//Session - session
func (c Clients) Session(profile string) *session.Session {
	if c.session != nil {
		return c.session
	}

	//Read AWS Config File

	// Initial credentials loaded from SDK's default credential chain. Such as
	// the environment, shared credentials (~/.aws/credentials), or EC2 Instance
	// Role. These credentials will be used to to make the STS Assume Role API.
	awsProfile := AWSProfile{}
	awsProfile = GetProfile(profile)
	//logLevel := aws.LogDebugWithHTTPBody

	config := aws.Config{
		Region: &awsProfile.Region,
		//LogLevel: &logLevel,
	}
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            config,
		Profile:           profile,
		SharedConfigState: session.SharedConfigEnable,
	}))

	c.session = sess
	return sess
}

//Config - it caches config for session
func (c Clients) Config(profile string) *aws.Config {

	// return no config for nil inputs
	if len(awsProfile.AccountID) == 0 || len(awsProfile.Region) == 0 || len(awsProfile.Role) == 0 {
		return nil
	}

	arn := fmt.Sprintf(
		"arn:aws:iam::%v:role/%v",
		awsProfile.AccountID,
		awsProfile.Role,
	)
	log.Printf("ARN: %s", arn)

	// include region in cache key otherwise concurrency errors
	key := fmt.Sprintf("%v::%v", awsProfile.Region, arn)

	log.Printf("Key: %s", key)

	// check for cached config
	if c.configs != nil && c.configs[key] != nil {
		return c.configs[key]
	}

	stsclient := sts.New(c.Session(profile))
	creds := credentials.NewChainCredentials([]credentials.Provider{

		&stscreds.AssumeRoleProvider{
			Client:  stsclient,
			RoleARN: arn,
		},
		&credentials.EnvProvider{},
		&ec2rolecreds.EC2RoleProvider{
			Client: ec2metadata.New(c.Session(profile)),
		},
		&credentials.SharedCredentialsProvider{
			Profile: profile,
		},
	})

	// new creds

	//creds = stscreds.NewCredentials(c.Session(), arn)
	if creds == nil {
	}

	// New config with Credentials
	config := aws.NewConfig().
		WithCredentials(creds).
		WithRegion(awsProfile.Region).
		WithMaxRetries(10)

	if c.configs == nil {
		c.configs = map[string]*aws.Config{}
	}

	c.configs[key] = config
	return config
}

// S3 - client
func (c *Clients) S3(profile string) s3iface.S3API {
	return s3.New(c.Session(profile), c.Config(profile))
}

// SQS - client
func (c *Clients) SQS(profile string) sqsiface.SQSAPI {
	return sqs.New(c.Session(profile), c.Config(profile))
}

// DynamoDB - client
func (c *Clients) DynamoDB(profile string) dynamodbiface.DynamoDBAPI {
	return dynamodb.New(c.Session(profile), c.Config(profile))
}
