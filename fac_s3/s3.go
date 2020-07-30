package facs3

import (
	"github.com/kataras/golog"
	facclients "github.com/wagner-aos/go-fast-aws-connections/fac_clients"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

var (
	err   error
	s3API s3iface.S3API
)

//Start - initializes S3 client
func Start(region, profile, endpoint string) {
	s3API = facclients.S3(region, profile, endpoint)
}

//ListBuckets - list all s3 available buckets
func ListBuckets() (*s3.ListBucketsOutput, error) {

	result, err := s3API.ListBuckets(nil)
	if err != nil {
		golog.Errorf("[fac_s3]-Error: %s", err)
		return nil, err
	}

	for _, b := range result.Buckets {
		golog.Errorf("[fac_s3]-* %s created on %s",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	return result, nil

}

//PrintBuckets - print all s3 available buckets
func PrintBuckets() error {

	result, err := s3API.ListBuckets(nil)
	if err != nil {
		golog.Errorf("[fac_s3]-Error: %s", err)
	}

	for _, b := range result.Buckets {
		golog.Infof("[fac_s3]-* %s created on %s",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	return err
}
