package role

type Role struct {
	code string
}

func (r *Role) Code() string {
	return r.code
}

func (r *Role) SetCode(code string) {
	r.code = code
}
