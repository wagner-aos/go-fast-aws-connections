package facsqs

import (
	"asappay-payment-processor/awscredentials"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

var (
	err    error
	sqsAPI sqsiface.SQSAPI
)

func init() {
	c := &awscredentials.Clients{}
	sqsAPI = c.SQS("asappay-Dev")
}

//SendMessage - it sends message to any SQS Queue
func SendMessage(queueName string, message string) {

	qURL := getQueueURL(queueName)
	fmt.Println(qURL)

	//DelaySeconds: aws.Int64(10),
	result, err := sqsAPI.SendMessage(&sqs.SendMessageInput{
		MessageGroupId:         aws.String("POS"),
		MessageDeduplicationId: aws.String("1234"),
		MessageBody:            aws.String(message),
		QueueUrl:               qURL,
	})

	if err != nil {
		fmt.Printf("Error sending message to queue: %s , %s ", queueName, err)
		return
	}

	fmt.Println("Success", *result.MessageId)
}

//getQueueURL - get queue entire URL in order to send messages to SQS.
func getQueueURL(queueName string) *string {

	output, err := sqsAPI.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		fmt.Println("Error recovering queueURL:", err)
	}
	return output.QueueUrl
}
