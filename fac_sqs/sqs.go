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

//Start - initializes SQS client
func Start(profile string) {
	sqsAPI = facclients.SQS(profile)
}

//SendMessage - it sends message to any SQS Queue
func SendMessage(messageInput *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {

	queueName := messageInput.QueueUrl
	queueURL := GetQueueURL(queueName)
	print(queueURL)
	messageInput.SetQueueUrl(*queueURL)

	//DelaySeconds: aws.Int64(10),
	result, err := sqsAPI.SendMessage(messageInput)

	if err != nil {
		fmt.Printf("Error sending message to queue: %s , %s ", *queueName, err)
		return nil, err
	}

	fmt.Println("Success", *result.MessageId)
	return result, nil
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

//GetQueueURL - get queue entire URL in order to send messages to SQS.
func GetQueueURL(queueName *string) *string {

	output, err := sqsAPI.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: queueName,
	})
	if err != nil {
		fmt.Println("Error recovering queueURL:", err)
	}
	return output.QueueUrl
}
