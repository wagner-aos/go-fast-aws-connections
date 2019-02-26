package main

import (
	"go-fast-aws-connections/fac_dynamodb"
	"go-fast-aws-connections/fac_s3"
	"go-fast-aws-connections/fac_sqs"
)

func main() {

	profile := "asappay-Mgmt"

	//profile
	facs3.Init(profile)
	facs3.ListBuckets()

	facsqs.Init(profile)
	facsqs.ListQueues()

	facdynamodb.Init(profile)
	facdynamodb.ListTables()

}
