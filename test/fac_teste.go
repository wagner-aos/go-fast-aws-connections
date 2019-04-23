package teste

import (
	"github.com/wagner-aos/go-fast-aws-connections/fac_dynamodb"
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
