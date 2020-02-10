package facdynamodb

import (
	"os"

	"github.com/kataras/golog"
	facclients "github.com/wagner-aos/go-fast-aws-connections/fac_clients"

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
func Start(region string, profile string) {
	dynamodbAPI = facclients.DynamoDB(region, profile)
}

//PutItem - It inputs and item into a dynamo table.
func PutItem(tableName string, object interface{}) (*dynamodb.PutItemOutput, error) {
	//Struct to DynamoItem
	dynamodbAttributes, err := dynamodbattribute.MarshalMap(object)
	if err != nil {
		golog.Error("[fac_dynamo]-Got error marshalling dynamo attributes map:")
		golog.Error(err.Error())
		os.Exit(1)
	}

	result, err := dynamodbAPI.PutItem(&dynamodb.PutItemInput{
		Item:         dynamodbAttributes,
		TableName:    aws.String(tableName),
		ReturnValues: aws.String("ALL_OLD"),
	})
	if err != nil {
		golog.Errorf("[fac_dynamo]-Error when put item into DynamoDB: %s , %s ", tableName, err)
		return result, err
	}

	//jsonOutPut, _ := json.Marshal(result.Attributes)
	golog.Infof("[fac_dynamo]-Success: %s", result.String())

	return result, nil
}

//Query - It queries items in a dynamodb table.
func Query(queryInput *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	result, err := dynamodbAPI.Query(queryInput)
	if err != nil {
		golog.Errorf("[fac_dynamo]-Error: %s", err)
		golog.Warn("[fac_dynamo]-Item not found!")
	}
	return result, nil
}

//Scan - It scans items in a dynamodb table.
func Scan(scanInput *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {

	result, err := dynamodbAPI.Scan(scanInput)
	if err != nil {
		golog.Errorf("[fac_dynamo]-Error: %s", err)
		golog.Warn("[fac_dynamo]-Item not found!")
	}
	return result, nil
}

//ListTables - list all dynamodb available tables
func ListTables() {

	result, err := dynamodbAPI.ListTables(nil)
	if err != nil {
		golog.Errorf("[fac_dynamo]-Error: %x", err)
	}

	for _, t := range result.TableNames {
		golog.Infof("[fac_dynamo]-* %s ", aws.StringValue(t))
	}

}
