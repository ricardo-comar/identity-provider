package model

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type ExecutionContext struct {
	Ctx       context.Context
	Cfg       aws.Config
	EventID   string
	ProcessID string
	MessageID string
	Body      string
}

func NewExecutionContext(Ctx context.Context, Cfg aws.Config, EventID string, ProcessID string, MessageID string, Body string) *ExecutionContext {

	execCtx := &ExecutionContext{}
	execCtx.Ctx = Ctx
	execCtx.Cfg = Cfg
	execCtx.EventID = EventID
	execCtx.ProcessID = ProcessID
	execCtx.MessageID = MessageID
	execCtx.Body = Body

	return execCtx
}
