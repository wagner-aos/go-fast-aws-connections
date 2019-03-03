package test

import (
	"go-fast-aws-connections/fac_sqs"
	"testing"
)

var profile = "asappay-Dev"

func TestListQueues(t *testing.T) {
	facsqs.Start(profile)
	facsqs.ListQueues()
}

func TestSendMessage(t *testing.T) {
	facsqs.Start(profile)
	facsqs.SendMessage("fac_sqs", "{\"message\":\"TEST\"}")
}
