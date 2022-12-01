package gateway

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/uuid"
	"github.com/ricardo-comar/identity-provider/core"
	"github.com/ricardo-comar/identity-provider/model"
)

func SendMessage(core *core.Core, ctx *model.ExecutionContext, message interface{}) (*string, error) {

	body, _ := json.Marshal(message)

	svc := sqs.NewFromConfig(ctx.Cfg)

	localendpoint, found := os.LookupEnv("LOCALSTACK_HOSTNAME")
	if found {
		svc = sqs.New(sqs.Options{Credentials: ctx.Cfg.Credentials, EndpointResolver: sqs.EndpointResolverFromURL("http://" + localendpoint + ":" + os.Getenv("EDGE_PORT"))})
		ctx.EventID = uuid.NewString()
	}

	res, err := svc.SendMessage(ctx.Ctx, &sqs.SendMessageInput{
		MessageAttributes: map[string]types.MessageAttributeValue{
			"EventID": {
				DataType:    aws.String("String"),
				StringValue: aws.String(ctx.EventID),
			},
		}, MessageBody: aws.String(string(body)),
		QueueUrl: &core.Config.EmployeeQueue,
	})

	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res.MessageId, nil
}
