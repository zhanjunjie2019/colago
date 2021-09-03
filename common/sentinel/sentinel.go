package sentinel

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/hotspot"
)

func init() {
	err := ioc.InjectSimpleBean(new(Sentinel))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
}

type Sentinel struct {
	ready               bool
	circuitbreakerRules []*circuitbreaker.Rule
	hotspotRule         []*hotspot.Rule
}

func (s *Sentinel) New() ioc.AbsBean {
	s.Init()
	return s
}

func (s *Sentinel) Init() {
	if !s.ready {
		err := sentinel.InitDefault()
		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}
		s.ready = true
	}
}

func (s *Sentinel) AppendCircuitbreakerRules(roles ...*circuitbreaker.Rule) {
	if s.circuitbreakerRules == nil {
		s.circuitbreakerRules = make([]*circuitbreaker.Rule, 0)
	}
	if roles != nil {
		for _, role := range roles {
			s.circuitbreakerRules = append(s.circuitbreakerRules, role)
		}
	}
}

func (s *Sentinel) AppendHotspotRules(roles ...*hotspot.Rule) {
	if s.hotspotRule == nil {
		s.hotspotRule = make([]*hotspot.Rule, 0)
	}
	if roles != nil {
		for _, role := range roles {
			s.hotspotRule = append(s.hotspotRule, role)
		}
	}
}

func (s *Sentinel) LoadRules() error {
	s.Init()
	_, err := circuitbreaker.LoadRules(s.circuitbreakerRules)
	if err != nil {
		_, err = hotspot.LoadRules(s.hotspotRule)
	}
	return err
}

func (s *Sentinel) Entry(
	resource string,
	tryFn func() (interface{}, error),
	args ...interface{},
) (interface{}, error) {
	e, b := sentinel.Entry(resource, sentinel.WithArgs(args))
	if b != nil {
		panic(NewSentinelError(resource))
	} else {
		defer func() {
			e.Exit()
		}()
		return tryFn()
	}
}

func NewSentinelError(resource string) *SentinelError {
	return &SentinelError{
		resource: resource,
	}
}

type SentinelError struct {
	resource string
}

func (err SentinelError) Error() string {
	return "SentinelError : " + err.resource + " is Fused!"
}
