package service

import (
	"context"

	"github.com/galaxy-future/schedulx/api/types"
)

type ActionSvc interface {
	ExecAct(ctx context.Context, args interface{}, act types.Action) (interface{}, error)
	entryLog(ctx context.Context, act string, req interface{})
	exitLog(ctx context.Context, act string, req, resp interface{}, err error)
}

var _ ActionSvc = &NodeActSvc{}
var _ ActionSvc = &BridgXSvc{}
var _ ActionSvc = &ScheduleSvc{}
var _ ActionSvc = &TemplateSvc{}
