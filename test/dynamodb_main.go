package main

import (
	"fmt"
	"log"

	"github.com/wagner-aos/go-fast-aws-connections/fac_dynamodb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	query()
}

func query() {
	profile := "asappay-Dev"

	facdynamodb.Start(profile)

	posID := "00000002"
	merchantID := "000000000000001"

	queryInput := &dynamodb.QueryInput{
		TableName: aws.String("PaymentProcessorProduct"),
		IndexName: aws.String("posID-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			"posID": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(posID),
					},
				},
			},
			"merchantID": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String(merchantID),
					},
				},
			},
		},
	}

	result, err := facdynamodb.Query(queryInput)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Success:")
	fmt.Print(result.String())

}

func scan() {

	// profile := "asappay-Dev"

	// facdynamodb.Start(profile)

	// result, err := facdynamodb.Scan(&dynamodb.ScanInput{
	// 	TableName: aws.String("PaymentProcessorProduct"),
	// })
	// if err != nil {
	// 	log.Println(err)
	// }

	// for _, v := range result.Items {
	// 	//j, _ := json.Marshal(v)
	// 	print(string(v))
	// }

}
