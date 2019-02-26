package main

import (
	"asappay-payment-processor/awscredentials"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
)

func main() {

	c := awscredentials.Clients{}

	//profile "asappay-Dev"
	api := c.S3("asappay-Dev")

	result, err := api.ListBuckets(nil)
	if err != nil {
		fmt.Printf("Error: %x", err)
	}

	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
