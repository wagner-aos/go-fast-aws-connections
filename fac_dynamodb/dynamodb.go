package facdynamodb

import (
	"fmt"
	"go-fast-aws-connections/fac_clients"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	err         error
	dynamodbAPI dynamodbiface.DynamoDBAPI
)

//Start - initializes S3 client
func Start(profile string) {
	dynamodbAPI = facclients.DynamoDB("asappay-Dev")
}

//ListTables - list all dynamodb available tables
func ListTables() {

	result, err := dynamodbAPI.ListTables(nil)
	if err != nil {
		fmt.Printf("Error: %x", err)
	}

	for _, t := range result.TableNames {
		fmt.Printf("* %s \n",
			aws.StringValue(t))
	}

}
