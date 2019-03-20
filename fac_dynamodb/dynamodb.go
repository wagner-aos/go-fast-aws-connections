package facdynamodb

import (
	"fmt"
	"go-fast-aws-connections/fac_clients"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

var (
	err         error
	dynamodbAPI dynamodbiface.DynamoDBAPI
)

//Start - initializes S3 client
func Start(profile string) {
	dynamodbAPI = facclients.DynamoDB(profile)
}

//Query - It queries items in a dynamodb table.
func Query(queryInput *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {

	result, err := dynamodbAPI.Query(queryInput)
	if err != nil {
		log.Printf("Error: %s", err)
		log.Println("Item not found!")
	}
	return result, nil
}

//Scan - It scans items in a dynamodb table.
func Scan(scanInput *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {

	result, err := dynamodbAPI.Scan(scanInput)
	if err != nil {
		log.Printf("Error: %s", err)
		log.Println("Item not found!")
	}
	return result, nil
}

//ListTables - list all dynamodb available tables
func ListTables() {

	result, err := dynamodbAPI.ListTables(nil)
	if err != nil {
		log.Printf("Error: %s", err)
	}

	for _, t := range result.TableNames {
		fmt.Printf("* %s \n",
			aws.StringValue(t))
	}

}
