package opentelemetry

import (
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/common/protoactor"
)

func init() {
	ioc.AppendInjection(func() *OpentelemetryFilter {
		return &OpentelemetryFilter{}
	})
}

type OpentelemetryFilter struct {
	next protoactor.ActorClientFilter
}

func (o *OpentelemetryFilter) SetNext(filter protoactor.ActorClientFilter) {
	o.next = filter
}

func (o OpentelemetryFilter) Filter(clientActionArgs protoactor.ClientActionArgs) (rs interface{}, err error) {

	return
}
