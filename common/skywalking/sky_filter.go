package skywalking

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"e.coding.net/double-j/ego/colago/common/protoactor"
	"fmt"
)

func init() {
	err := ioc.InjectSimpleBean(new(SkyFilter))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

func SkyFilterFactory() protoactor.ActorClientFilter {
	bean, err := ioc.GetBean("skywalking.SkyFilter")
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
	chain := bean.(*SkyFilter)
	return chain
}

type SkyFilter struct {
	next protoactor.ActorClientFilter
}

func (s *SkyFilter) New() ioc.AbsBean {
	return s
}

func (s *SkyFilter) SetNext(filter protoactor.ActorClientFilter) {
	s.next = filter
}

func (s *SkyFilter) Filter(clientActionArgs protoactor.ClientActionArgs) (rs interface{}, err error) {
	parentSpan, err := GetSpanByCtx(clientActionArgs.Ctx)
	if err != nil {
		return nil, err
	}
	span, err := parentSpan.CreateExitSpan(clientActionArgs.OperationName, clientActionArgs.Peer, func(headerKey, headerValue string) error {
		return nil
	})
	defer func() {
		span.End(err)
	}()
	if err != nil {
		return nil, err
	}
	return s.next.Filter(clientActionArgs)
}
