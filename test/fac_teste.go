package teste

import (
	"go-fast-aws-connections/fac_dynamodb"
	"go-fast-aws-connections/fac_s3"
	"go-fast-aws-connections/fac_sqs"
)

func test() {

	profile := "asappay-Dev"

	//profile
	facs3.Start(profile)
	facs3.ListBuckets()

	facsqs.Start(profile)
	facsqs.ListQueues()

	facdynamodb.Start(profile)
	facdynamodb.ListTables()

}
