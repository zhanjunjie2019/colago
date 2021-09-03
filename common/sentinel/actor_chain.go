package sentinel

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(SentinuelActorChain))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func SentinuelActorChainFactory() protoactor.ActorClientFilter {
	bean, err := ioc.GetBean("sentinel.SentinuelActorChain")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	chain := bean.(*SentinuelActorChain)
	return chain
}

type SentinuelActorChain struct {
	next protoactor.ActorClientFilter
	Sent *Sentinel `ij:"sentinel.Sentinel"`
}

func (s *SentinuelActorChain) New() ioc.AbsBean {
	return s
}

func (s *SentinuelActorChain) SetNext(filter protoactor.ActorClientFilter) {
	s.next = filter
}

func (s *SentinuelActorChain) Filter(clientActionArgs protoactor.ClientActionArgs) (rs interface{}, err error) {
	return s.Sent.Entry(
		clientActionArgs.Resource,
		func() (interface{}, error) {
			return s.next.Filter(clientActionArgs)
		},
	)
}
