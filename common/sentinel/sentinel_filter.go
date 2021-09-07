package sentinel

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(SentinelFilter))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func SentinulFilterFactory() protoactor.ActorClientFilter {
	bean, err := ioc.GetBean("sentinel.SentinelFilter")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	chain := bean.(*SentinelFilter)
	return chain
}

type SentinelFilter struct {
	next protoactor.ActorClientFilter
	Sent *Sentinel `ij:"sentinel.Sentinel"`
}

func (s *SentinelFilter) New() ioc.AbsBean {
	return s
}

func (s *SentinelFilter) SetNext(filter protoactor.ActorClientFilter) {
	s.next = filter
}

func (s *SentinelFilter) Filter(clientActionArgs protoactor.ClientActionArgs) (rs interface{}, err error) {
	return s.Sent.Entry(
		clientActionArgs.OperationName+"."+clientActionArgs.Peer,
		func() (interface{}, error) {
			return s.next.Filter(clientActionArgs)
		},
	)
}
