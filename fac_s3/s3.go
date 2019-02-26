package facs3

import (
	"fmt"
	"go-fast-aws-connections/fac_clients"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

var (
	err   error
	s3API s3iface.S3API
)

//Init - initializes S3 client
func Init(profile string) {
	s3API = facclients.S3("asappay-Dev")
}

//ListBuckets - list all s3 available buckets
func ListBuckets() {

	result, err := s3API.ListBuckets(nil)
	if err != nil {
		fmt.Printf("Error: %x", err)
	}

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

}
