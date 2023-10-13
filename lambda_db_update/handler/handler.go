package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/google/uuid"
	"github.com/ricardo-comar/identity-provider/db_update/gateway"
	"github.com/ricardo-comar/identity-provider/lib_common/model"

	"encoding/json"
)

func main() {
	lambda.Start(handleMessages)
}

func handleMessages(ctx context.Context, sqsEvent events.SQSEvent) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("AWS_REGION")
		return nil
	})
	if err != nil {
		return err
	}

	processID := uuid.New().String()
	inicioProc := time.Now()
	log.Printf("Processo iniciado: %s", processID)

	for _, message := range sqsEvent.Records {
		eventID := message.MessageAttributes["EventID"].StringValue

		execCtx := model.NewExecutionContext(ctx, cfg, *eventID, processID, message.MessageId, message.Body)

		log.Printf("Iniciando - evento %s , processo %s e mensagem %s", execCtx.EventID, execCtx.ProcessID, execCtx.MessageID)
		inicioMsg := time.Now()

		handleMessage(execCtx, message.Body)

		log.Printf("Finalizando - mensagem %s em %dms", execCtx.MessageID, time.Since(inicioMsg).Milliseconds())
	}

	log.Printf("Finalizando - processo %s em %dms", processID, time.Since(inicioProc).Milliseconds())
	return nil
}

func handleMessage(ctx *model.ExecutionContext, msg string) (interface{}, error) {

	data := model.EmployeeMessage{}
	json.Unmarshal([]byte(msg), &data)

	log.Println("Saving Employee")
	gateway.SaveEmployee(ctx.Cfg, &data)

	return data, nil

}
