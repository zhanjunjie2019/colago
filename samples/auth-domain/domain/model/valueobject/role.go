package valueobject

import "e.coding.net/double-j/ego/colago/common/ioc"

func init() {
	_ = ioc.InjectPrototypeBean(new(Role))
}

type Role struct {
	code string
}

func (r *Role) New() ioc.AbsBean {
	return r
}

func (r *Role) Code() string {
	return r.code
}

func (r *Role) SetCode(code string) {
	r.code = code
}
