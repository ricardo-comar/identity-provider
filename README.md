
# Identity Provider

Welcome to my instructional application, witch helped me to learn GoLang and Terraform :smile: 

In this repository you will find a simple but helpful example of three types of trigger a Lambda (eventbridge, sqs and http request). All together glued by terraform to proccess and persist data, to be retrieven later using an API interface. 

## Blueprint
![](doc/aws-blueprint.png)

## Components

### Registration Data Lambda

Lambda responsible to query for outside data, in this example from a mock origin ([mockaroo.com]()), split in individual records, and post all of them on SQS queue.

### DB Update Lambda

Lambda responsible to read the received data and persist on a DynamoDB table

### REST API Lambda

Lambda responsible to expose the persisted data from DybamoDB

### API Gateway

Component responsible to expose the Lambda to outside of the VPC

### DynamoDB

Component responsible to retain the data, using a TTL (time to live) column to automaticaly remove the obsolete data.

### Bucket S3, Glaciar and Athena :construction:

Will be available upon issue #1 is implemented

- [ ] https://github.com/ricardo-comar/identity-provider/issues/1



## Workspace configuration

To execute this application, I recomend you to install the following resources:

- VS Code
- Plugins - Go, Terraform
- Go (golang)
- AWS CLI
- Terraform
- Localstack
- tflocal
- docker
- docker-compose

### Execution

First you need to compile the project using the Makefile script:

> make package

Second, start the **localstack** infrastructure to emulate AWS environment:

> cd localstack
> docker-compose up

In another terminal configure your AWS CLI, informing same access and secret keys from [docker-compose.yml](localstack/docker-compose.yaml):

> aws configure

Finally, run Terraform commands:

> tflocal init
> tflocal apply -auto-approve

After creating all resources, a scheduled trigger will start the first Lambda after 2 minutes, quering the data and splitting into messages to be persisted on DynamoDB.

To query for the API id, run this following command:

> aws --endpoint-url=http://localhost:4566 apigateway get-rest-apis
```
{
    "items": [
        {
            "id": "uw6qnzhus8",
            "name": "idp_api",
            "createdDate": "2022-11-25T14:22:41-03:00",
            "version": "V1",
...
    ]
}
```

To call the API, use this URL below with _curl_ changing the ID retrieved: 

> curl http://localhost:4566/restapis/"id"/v1/\_user_request_/employees

More info: https://docs.localstack.cloud/aws/apigatewayv2/



## References
