package model

type BizError struct {
	Message string
}

func (err BizError) Error() string {
	return "BIZ WARN : " + err.Message
}
