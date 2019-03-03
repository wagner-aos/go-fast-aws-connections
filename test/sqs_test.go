package sqstest

import (
	"go-fast-aws-connections/fac_sqs"
	"testing"
)

func TestSendMessage(t *testing.T) {

	profile := "asappay-Mgmt"

	facsqs.Start(profile)
	facsqs.ListQueues()

}
