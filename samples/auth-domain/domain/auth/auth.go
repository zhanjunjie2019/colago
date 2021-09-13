package auth

type Auth struct {
	code string
}

func (a *Auth) Code() string {
	return a.code
}

func (a *Auth) SetCode(code string) {
	a.code = code
}
