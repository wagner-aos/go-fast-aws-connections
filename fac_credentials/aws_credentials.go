package awscredentials

//package main

//https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/configuring-sdk.html
//https://github.com/aws/aws-sdk-go
//https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sessions.html

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/kataras/golog"
)

var awsProfile AWSProfile

//Clients - It store config of each session
type Clients struct {
	session  *session.Session
	configs  map[string]*aws.Config
	profiles map[string]*string
}

//Session - session
func (c Clients) Session() *session.Session {
	if c.session != nil {
		return c.session
	}

	//Read AWS Config File

	// Initial credentials loaded from SDK's default credential chain. Such as
	// the environment, shared credentials (~/.aws/credentials), or EC2 Instance
	// Role. These credentials will be used to to make the STS Assume Role API.

	logLevel := aws.LogDebugWithHTTPBody

	sess := session.Must(session.NewSession(
		&aws.Config{
			//Region: &awsProfile.Region,
			LogLevel: &logLevel,
		}))

	c.session = sess
	return sess
}

//SessionWithProfile - session
func (c Clients) SessionWithProfile(profile string) *session.Session {
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
	golog.Infof("ARN: %s", arn)

	// include region in cache key otherwise concurrency errors
	key := fmt.Sprintf("%v::%v", awsProfile.Region, arn)

	golog.Infof("Key: %s", key)

	// check for cached config
	if c.configs != nil && c.configs[key] != nil {
		return c.configs[key]
	}

	stsclient := sts.New(c.SessionWithProfile(profile))
	creds := credentials.NewChainCredentials([]credentials.Provider{

		&stscreds.AssumeRoleProvider{
			Client:  stsclient,
			RoleARN: arn,
		},
		&credentials.EnvProvider{},
		&ec2rolecreds.EC2RoleProvider{
			Client: ec2metadata.New(c.Session()),
		},
		&credentials.SharedCredentialsProvider{
			Profile: profile,
		},
	})

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
