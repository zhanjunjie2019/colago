package domain

import "e.coding.net/double-j/ego/colago/common/ioc"

var domainFactory = new(DomainFactory)

func GetDomainFactory() *DomainFactory {
	return domainFactory
}

type DomainFactory struct {
}

func (d *DomainFactory) Create(clazzName string) (ioc.AbsBean, error) {
	return ioc.GetBean(clazzName)
}
