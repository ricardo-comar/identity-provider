package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/ricardo-comar/identity-provider/model"
	"github.com/ricardo-comar/identity-provider/service"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/google/uuid"
)

var cfg aws.Config

func init() {
	cfg, _ = config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("AWS_REGION")
		return nil
	})
}

func main() {
	lambda.Start(eventHandler)
}

func eventHandler(ctx context.Context, event events.CloudWatchEvent) (events.CloudWatchEvent, error) {

	eventjson, err := json.Marshal(event)
	println("event= ", string(eventjson))

	execCtx := model.NewExecutionContext(ctx, cfg, event.ID, uuid.New().String(), event.Source, string(event.Detail))

	log.Printf("Iniciando - evento %s e processo %s", execCtx.EventID, execCtx.ProcessID)
	inicio := time.Now()

	_, err = service.EmployeeService(execCtx)
	if err != nil {
		panic(err)
	}

	log.Printf("Finalizando - processo %s em %dms", execCtx.ProcessID, time.Now().Sub(inicio))
	return event, nil
}
