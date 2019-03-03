package test

import (
	"go-fast-aws-connections/fac_dynamodb"
	"go-fast-aws-connections/fac_s3"
	"go-fast-aws-connections/fac_sqs"
)

func testes() {

	profile := "asappay-Mgmt"

	//profile
	facs3.Init(profile)
	facs3.ListBuckets()

	facsqs.Start(profile)
	facsqs.ListQueues()

	facdynamodb.Init(profile)
	facdynamodb.ListTables()

}
