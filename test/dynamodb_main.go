package test

import (
	"github.com/kataras/golog"
	facdynamodb "github.com/wagner-aos/go-fast-aws-connections/fac_dynamodb"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	query()
}

func query() {
	profile := "asappay-dev"
	region := "us-east-1"

	facdynamodb.Start(region, profile, "")

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
		golog.Error(err)
	}

	if len(result.String()) > 0 {
		golog.Infof("success: %s", result.String())
	}

}

func scan() {

	// profile := "asappay-Dev"

	// facdynamodb.Start(profile)

	// result, err := facdynamodb.Scan(&dynamodb.ScanInput{
	// 	TableName: aws.String("PaymentProcessorProduct"),
	// })
	// if err != nil {
	// 	golog.Info(err)
	// }

	// for _, v := range result.Items {
	// 	//j, _ := json.Marshal(v)
	// 	print(string(v))
	// }

}
