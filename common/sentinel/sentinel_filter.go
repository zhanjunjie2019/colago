package sentinel

import (
	"github.com/zhanjunjie2019/colago/common/ioc"
	"github.com/zhanjunjie2019/colago/common/protoactor"
)

func init() {
	ioc.AppendInjection(func(sent *Sentinel) *SentinelFilter {
		return &SentinelFilter{
			sent: sent,
		}
	})
}

type SentinelFilter struct {
	sent *Sentinel
	next protoactor.ActorClientFilter
}

func (s *SentinelFilter) SetNext(filter protoactor.ActorClientFilter) {
	s.next = filter
}

func (s *SentinelFilter) Filter(clientActionArgs protoactor.ClientActionArgs) (rs interface{}, err error) {
	rs, err = s.sent.Entry(
		clientActionArgs.OperationName,
		func() (interface{}, error) {
			return s.next.Filter(clientActionArgs)
		},
	)
	return
}
