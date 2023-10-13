module github.com/ricardo-comar/identity-provider/rest_api

go 1.21

require (
	github.com/aws/aws-sdk-go-v2/credentials v1.13.0 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.12.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.1.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.4.19 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.3.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodbstreams v1.13.23 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.9.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.7.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.9.19 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.11.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.13.8 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.17.2 // indirect
	github.com/aws/smithy-go v1.15.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
)

require (
	github.com/aws/aws-lambda-go v1.34.1
	github.com/aws/aws-sdk-go-v2 v1.21.2
	github.com/aws/aws-sdk-go-v2/config v1.18.0
	github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue v1.10.3
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.17.4
	github.com/ricardo-comar/identity-provider/lib_common v0.0.0-00010101000000-000000000000
)

replace github.com/ricardo-comar/identity-provider/lib_common => ../lib_common
