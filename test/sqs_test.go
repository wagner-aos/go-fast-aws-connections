package test

import (
	"testing"

	"github.com/wagner-aos/go-fast-aws-connections/fac_sqs"
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
