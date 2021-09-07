package skywalking

import (
	"github.com/SkyAPM/go2sky"
	"golang.org/x/net/context"
	"time"
)

type skyspan struct {
	ctx            context.Context
	skyWalkingSpan go2sky.Span
}

func (s *skyspan) Ctx() context.Context {
	return s.ctx
}

func (s *skyspan) TraceID() string {
	return go2sky.TraceID(s.ctx)
}

func (s *skyspan) CreateLocalSpan(operationName string) (Span, error) {
	nextSpan, ctx, err := tracer.CreateLocalSpan(s.ctx)
	if err != nil {
		return nil, err
	}
	nextSpan.SetOperationName(operationName)
	span := &skyspan{
		skyWalkingSpan: nextSpan,
	}
	ctx = context.WithValue(ctx, spanKeyInstance, span)
	span.ctx = ctx
	return span, nil
}

func (s *skyspan) CreateExitSpan(operationName string, remoteAddr string, injector func(headerKey, headerValue string) error) (Span, error) {
	nextSpan, err := tracer.CreateExitSpan(s.ctx, operationName, remoteAddr, injector)
	if err != nil {
		return nil, err
	}
	return &skyspan{
		ctx:            s.ctx,
		skyWalkingSpan: nextSpan,
	}, nil
}

func (s *skyspan) End(errs ...error) {
	var errsInfo []string
	for _, v := range errs {
		if v != nil {
			errsInfo = append(errsInfo, v.Error())
		}
	}
	if len(errsInfo) > 0 {
		s.skyWalkingSpan.Error(time.Now(), errsInfo...)
	}
	s.skyWalkingSpan.End()
}
