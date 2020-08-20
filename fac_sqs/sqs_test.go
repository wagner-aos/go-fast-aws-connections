package facsqs

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/stretchr/testify/assert"
)

var profile = "asappay-dev"
var region = "us-east-1"

func TestListQueues(t *testing.T) {
	Start(region, profile, "")
	//Start(region, profile, "http://localhost:9324")
	queues := ListQueues()
	t.Logf("%v", queues)
	assert.NotEmpty(t, queues)
}

func TestGetQueueUrl(t *testing.T) {
	Start(region, profile, "")
	queueURL, err := GetQueueURL("fac_sqs")
	t.Logf(" QueueUrl = %s", *queueURL)
	t.Logf(" ERROR = %s", err)
	assert.Nil(t, err)
	assert.NotEmpty(t, queueURL)
}

func TestSendMessage(t *testing.T) {
	Start(region, profile, "")
	//sqsOutPut := &sqs.SendMessageOutput{}
	output, err := SendMessage("fac_sqs", "{\"message\":\"TEST\"}")

	t.Logf(" Message = %s", *output.MessageId)

	t.Logf(" ERROR = %s", err)
	assert.Nil(t, err)
}

func TestReceiveMessage(t *testing.T) {
	Start(region, profile, "")

	output, err := ReceiveMessage("fac_sqs")

	t.Logf(" Message = %s", *output.Messages[0].ReceiptHandle)
	t.Logf(" ERROR = %s", err)
	assert.Nil(t, err)
}

func TestReceiveMessageInput(t *testing.T) {
	Start(region, profile, "")

	queueURL, err := GetQueueURL("fac_sqs")
	t.Logf(" Queue URL: = %s", *queueURL)

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

	req := *receiveMessageInput
	t.Logf(" REQ = %v", req)

	output, err := ReceiveMessageInput(receiveMessageInput)

	t.Logf(" Message = %s", *output.Messages[0].ReceiptHandle)

	if len(output.Messages[0].Attributes) > 0 {
		t.Logf(" Message = %s", *output.Messages[0].ReceiptHandle)
	}
	t.Logf(" ERROR = %s", err)
	assert.Nil(t, err)
}
