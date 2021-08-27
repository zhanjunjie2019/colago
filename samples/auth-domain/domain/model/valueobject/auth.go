package valueobject

import "e.coding.net/double-j/ego/colago/common/ioc"

func init() {
	_ = ioc.InjectPrototypeBean(new(Auth))
}

type Auth struct {
	code string
}

func (a *Auth) New() ioc.AbsBean {
	return a
}

func (a *Auth) Code() string {
	return a.code
}

func (a *Auth) SetCode(code string) {
	a.code = code
}
