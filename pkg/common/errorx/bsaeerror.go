package errorx

const DefaultCode = 0

type CodeError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (e CodeError) Error() string {
	return e.Message
}

func NewCodeError(code int, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

func NewDefaultCodeError(msg string) error {
	return NewCodeError(DefaultCode, msg)
}
