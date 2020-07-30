package facsqs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var profile = "asappay-dev"
var region = "us-east-1"

func TestListQueues(t *testing.T) {
	Start(region, profile, "")
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
}

func TestSendMessage(t *testing.T) {
	Start(region, profile, "")
	//sqsOutPut := &sqs.SendMessageOutput{}
	_, err = SendMessage("fac_sqs", "{\"message\":\"TEST\"}")

	t.Logf(" ERROR = %s", err)
	assert.Nil(t, err)

}
