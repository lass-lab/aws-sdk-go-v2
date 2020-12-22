module github.com/aws/aws-sdk-go-v2/service/s3control

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v0.30.1-0.20201222223005-ee883de66531
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v0.3.3-0.20201222223005-ee883de66531
	github.com/aws/smithy-go v0.4.1-0.20201222001052-74df8ddd8c79
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/service/internal/s3shared => ../../service/internal/s3shared/
