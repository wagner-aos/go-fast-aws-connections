package main

import (
	"go-fast-aws-connections/fac_sqs"
)

func main() {

	profile := "asappay-Dev"

	facsqs.Start(profile)
	facsqs.ListQueues()

}
