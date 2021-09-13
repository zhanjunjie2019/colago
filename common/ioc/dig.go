package ioc

import "go.uber.org/dig"

var c *dig.Container

var constructors = make([]interface{}, 0)

var pullFactorys = make([]interface{}, 0)

func init() {
	c = dig.New()
}

func GetContainer() *dig.Container {
	return c
}

func AppendInjection(constructor interface{}) {
	constructors = append(constructors, constructor)
}

func AppendPullFactory(factory interface{}) {
	pullFactorys = append(pullFactorys, factory)
}

func BatchProvideFinal() {
	for _, constructor := range constructors {
		c.Provide(constructor)
	}
	for _, factory := range pullFactorys {
		c.Invoke(factory)
	}
}
