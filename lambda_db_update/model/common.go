package model

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

type ExecutionContext struct {
	Cfg       aws.Config
	EventID   string
	ProcessID string
	MessageID string
	Body      string
}

func NewExecutionContext(cfg aws.Config, EventID string, ProcessID string, MessageID, Body string) *ExecutionContext {

	ctx := &ExecutionContext{}
	ctx.Cfg = cfg
	ctx.EventID = EventID
	ctx.ProcessID = ProcessID
	ctx.MessageID = MessageID
	ctx.Body = Body

	return ctx
}
