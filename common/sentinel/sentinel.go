package sentinel

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/circuitbreaker"
	"github.com/alibaba/sentinel-golang/core/hotspot"
	"github.com/zhanjunjie2019/colago/common/ioc"
)

func init() {
	ioc.AppendInjection(func() *Sentinel {
		return new(Sentinel)
	})
}

type Sentinel struct {
	ready               bool
	circuitbreakerRules []*circuitbreaker.Rule
	hotspotRule         []*hotspot.Rule
}

func (s *Sentinel) init() {
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
	s.init()
	_, err := circuitbreaker.LoadRules(s.circuitbreakerRules)
	if err != nil {
		_, err = hotspot.LoadRules(s.hotspotRule)
	}
	return err
}

func (Sentinel) Entry(
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
