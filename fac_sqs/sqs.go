package facsqs

import (
	"fmt"

	"go-fast-aws-connections/fac_clients"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

var (
	err    error
	sqsAPI sqsiface.SQSAPI
)

//Init - initializes SQS client
func Init(profile string) {
	sqsAPI = facclients.SQS("asappay-Dev")
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

//ListQueues - list all available sqs queues
func ListQueues() {

	result, err := sqsAPI.ListQueues(nil)
	if err != nil {
		fmt.Printf("Error: %x", err)
	}

	for _, b := range result.QueueUrls {
		fmt.Printf("* %s \n",
			aws.StringValue(b))
	}

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
