package opentelemetry

import (
	"context"
	"github.com/zhanjunjie2019/colago/common/conf"
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/common/protoactor"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func init() {
	InitProvider(context.Background())
	ioc.AppendInjection(func() *OpentelemetryFilter {
		tracerName := conf.ConfigMap("otel.tracer.name")
		return &OpentelemetryFilter{
			tracer: otel.Tracer(tracerName),
		}
	})
}

type OpentelemetryFilter struct {
	tracer trace.Tracer
	next   protoactor.ActorClientFilter
}

func (o *OpentelemetryFilter) SetNext(filter protoactor.ActorClientFilter) {
	o.next = filter
}

func (o *OpentelemetryFilter) Filter(clientActionArgs protoactor.ClientActionArgs) (rs interface{}, err error) {
	ctx, span := o.tracer.Start(
		clientActionArgs.Ctx,
		clientActionArgs.Peer,
	)
	clientActionArgs.Ctx = ctx
	o.next.Filter(clientActionArgs)
	span.End()
	return
}
