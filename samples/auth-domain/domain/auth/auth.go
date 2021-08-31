package auth

import (
	"e.coding.net/double-j/ego/colago/common/ioc"
	"fmt"
)

func init() {
	err := ioc.InjectPrototypeBean(new(Auth))
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}
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
