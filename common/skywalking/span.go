package skywalking

import "golang.org/x/net/context"

type Span interface {
	Ctx() context.Context
	TraceID() string
	CreateLocalSpan(operationName string) (Span, error)
	CreateExitSpan(operationName string, remoteAddr string, injector func(headerKey, headerValue string) error) (Span, error)
	End(errs ...error)
}
