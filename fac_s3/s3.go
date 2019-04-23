package facs3

import (
	"fmt"

	"github.com/wagner-aos/go-fast-aws-connections/fac_clients"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

var (
	err   error
	s3API s3iface.S3API
)

//Start - initializes S3 client
func Start(profile string) {
	s3API = facclients.S3(profile)
}

//ListBuckets - list all s3 available buckets
func ListBuckets() (*s3.ListBucketsOutput, error) {

	result, err := s3API.ListBuckets(nil)
	if err != nil {
		fmt.Printf("Error: %x", err)
		return nil, err
	}

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	return result, nil

}

//PrintBuckets - print all s3 available buckets
func PrintBuckets() error {

	result, err := s3API.ListBuckets(nil)
	if err != nil {
		fmt.Printf("Error: %x", err)
	}

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	return err
}
