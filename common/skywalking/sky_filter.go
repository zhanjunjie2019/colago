package skywalking

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
)

func init() {
	ioc.AppendInjection(func() *SkyFilter {
		return new(SkyFilter)
	})
}

type SkyFilter struct {
	next protoactor.ActorClientFilter
}

func (s *SkyFilter) SetNext(filter protoactor.ActorClientFilter) {
	s.next = filter
}

func (s *SkyFilter) Filter(clientActionArgs protoactor.ClientActionArgs) (rs interface{}, err error) {
	parentSpan, err := GetSpanByCtx(clientActionArgs.Ctx)
	if err != nil {
		return nil, err
	}
	span, err := parentSpan.CreateExitSpan(clientActionArgs.OperationName, clientActionArgs.Peer, clientActionArgs.SetTraceId)
	defer func() {
		span.End(err)
	}()
	if err != nil {
		return nil, err
	}
	return s.next.Filter(clientActionArgs)
}
