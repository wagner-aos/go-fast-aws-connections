package test

import facsqs "github.com/wagner-aos/go-fast-aws-connections/fac_sqs"
import facs3 "github.com/wagner-aos/go-fast-aws-connections/fac_s3"
import facdynamodb "github.com/wagner-aos/go-fast-aws-connections/fac_dynamodb"

func test() {

	profile := "asappay-dev"
	region := "us-east-1"

	//profile
	facs3.Start(region, profile, "")
	facs3.ListBuckets()

	facsqs.Start(region, profile, "")
	facsqs.ListQueues()

	facdynamodb.Start(region, profile, "")
	facdynamodb.ListTables()

}
