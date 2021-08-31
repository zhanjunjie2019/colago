package role

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"fmt"
)

func init() {
	err := ioc.InjectPrototypeBean(new(Role))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
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
