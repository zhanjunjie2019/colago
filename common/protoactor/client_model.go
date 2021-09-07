package protoactor

import (
	"golang.org/x/net/context"
)

type ClientActionArgs struct {
	Ctx           context.Context
	OperationName string
	Peer          string
	SetTraceId    func(key, value string) error
	TryFn         func() (interface{}, error)
	CatchFn       func(interface{}) (interface{}, error)
}
