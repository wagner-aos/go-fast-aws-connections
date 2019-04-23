package facdynamodb

import (
	"fmt"
	"log"
	"os"

	"github.com/wagner-aos/go-fast-aws-connections/fac_clients"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
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

//PutItem - It inputs and item into a dynamo table.
func PutItem(tableName string, object interface{}) (*dynamodb.PutItemOutput, error) {
	//Struct to DynamoItem
	dynamodbAttributes, err := dynamodbattribute.MarshalMap(object)
	if err != nil {
		fmt.Println("Got error marshalling dynamo attributes map:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	result, err := dynamodbAPI.PutItem(&dynamodb.PutItemInput{
		Item:         dynamodbAttributes,
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_OLD"),
	})
	if err != nil {
		fmt.Printf("Error when put item into DynamoDB: %s , %s ", tableName, err)
		return result, err
	}

	//jsonOutPut, _ := json.Marshal(result.Attributes)
	fmt.Println("Success:")
	fmt.Print(result.String())

	return result, nil
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
