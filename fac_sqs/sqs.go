package facsqs

import (
	"github.com/kataras/golog"
	facclients "github.com/wagner-aos/go-fast-aws-connections/fac_clients"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

var (
	err    error
	sqsAPI sqsiface.SQSAPI
)

//Start - initializes SQS client
func Start(region, profile, endpoint string) {
	sqsAPI = facclients.SQS(region, profile, endpoint)
}

//SendMessageInputToQueueURL - it sends message input to any SQS Queue URL
func SendMessageInputToQueueURL(messageInput *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	return messageSender(messageInput)
}

//SendMessageInput - it sends message input to any SQS Queue
func SendMessageInput(messageInput *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	queueURL, err := GetQueueURL(*messageInput.QueueUrl)
	if err != nil {
		return nil, err
	}
	messageInput.SetQueueUrl(*queueURL)
	return messageSender(messageInput)
}

//SendMessage - it sends message to any SQS Queue
func SendMessage(queueName string, message string) (*sqs.SendMessageOutput, error) {

	queueURL, err := GetQueueURL(queueName)
	if err != nil {
		return nil, err
	}

	messageInput := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		//MessageGroupId: aws.String("GroupID"),
		QueueUrl: queueURL,
	}
	return messageSender(messageInput)
}

//SendMessageToQueueURL - it sends message to any SQS Queue URL
func SendMessageToQueueURL(queueURL string, message string) (*sqs.SendMessageOutput, error) {

	messageInput := &sqs.SendMessageInput{
		MessageBody: aws.String(message),
		//MessageGroupId: aws.String("GroupID"),
		QueueUrl: aws.String(queueURL),
	}

	return messageSender(messageInput)
}

func messageSender(messageInput *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	result, err := sqsAPI.SendMessage(messageInput)
	if err != nil {
		golog.Errorf("[fac_sqs]-Error sending message to queue: %s , %s ", *messageInput.QueueUrl, err)
		return nil, err
	}

	golog.Info("[fac_sqs]-Send Message OK.")
	golog.Infof("[fac_sqs]-MessageID: %s", *result.MessageId)
	return result, nil
}

//ListQueues - list all available sqs queues
func ListQueues() []*string {

	result, err := sqsAPI.ListQueues(nil)
	if err != nil {
		golog.Errorf("[fac_sqs]-Error listing queues: %s", err)
	}

	for _, b := range result.QueueUrls {
		golog.Infof("[fac_sqs]-* %s", aws.StringValue(b))
	}

	return result.QueueUrls
}

//ReceiveMessage - it receives message from any SQS Queue
func ReceiveMessage(queueName string) (*sqs.ReceiveMessageOutput, error) {

	queueURL, err := GetQueueURL(queueName)
	if err != nil {
		return nil, err
	}

	receiveMessageInput := &sqs.ReceiveMessageInput{
		AttributeNames: []*string{
			aws.String(sqs.MessageSystemAttributeNameSentTimestamp),
		},
		MessageAttributeNames: []*string{
			aws.String(sqs.QueueAttributeNameAll),
		},
		QueueUrl:            queueURL,
		MaxNumberOfMessages: aws.Int64(int64(1)),
		WaitTimeSeconds:     aws.Int64(int64(1)),
	}
	output, err := messageReceiver(receiveMessageInput)
	return output, err
}

//ReceiveMessageInput - it receives message input to any SQS Queue
func ReceiveMessageInput(messageInput *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	output, err := messageReceiver(messageInput)
	return output, err
}

func messageReceiver(receiveMessageInput *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
	output, err := sqsAPI.ReceiveMessage(receiveMessageInput)
	if err != nil {
		golog.Errorf("[fac_sqs] - %s", err)
		return nil, err
	}
	golog.Debug("[fac_sqs]-Receive Message OK.")
	golog.Debugf("[fac_sqs]-Message: %v", output.Messages)
	return output, nil
}

//DeleteMessageInput - it deletes messages
func DeleteMessageInput(deleteMessageInput *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	output, err := messageDelete(deleteMessageInput)
	return output, err
}

//DeleteMessageBatchInput - it deletes batch messages
func DeleteMessageBatchInput(deleteMessageBatchInput *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
	output, err := messageBatchDelete(deleteMessageBatchInput)
	return output, err
}

func messageDelete(deleteMessageInput *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
	output, err := sqsAPI.DeleteMessage(deleteMessageInput)
	if err != nil {
		golog.Errorf("[fac_sqs] - %s", err)
		return nil, err
	}
	golog.Debug("[fac_sqs]-Delete Message OK.")
	golog.Debugf("[fac_sqs]- %v", output.String())
	return output, nil
}

func messageBatchDelete(deleteMessageBatchInput *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
	output, err := sqsAPI.DeleteMessageBatch(deleteMessageBatchInput)
	if err != nil {
		golog.Errorf("[fac_sqs] - %s", err)
		return nil, err
	}
	golog.Debug("[fac_sqs]-Delete Batch Message OK.")
	golog.Debugf("[fac_sqs]- %v", output.String())
	return output, nil
}

//GetQueueURL - get queue entire URL in order to send messages to SQS.
func GetQueueURL(queueName string) (*string, error) {
	output, err := sqsAPI.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		golog.Errorf("[fac_sqs]-Error recovering queueURL: %s", err)
		return nil, err
	}
	return output.QueueUrl, nil
}
