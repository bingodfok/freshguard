package result

import (
	"time"
)

type Result struct {
	Code      int         `json:"code"`
	Msg       string      `json:"msg"`
	Timestamp int64       `json:"timestamp"`
	Data      interface{} `json:"data,omitempty"`
}

func NewResult(code int, msg string, data interface{}) *Result {
	return &Result{Code: code, Msg: msg, Timestamp: time.Now().UnixMilli(), Data: data}
}

func Success(data interface{}) *Result {
	return NewResult(200, "Success", data)
}

func CommonFail(msg string) *Result {
	return NewResult(400, msg, nil)
}

func Fail(code int, msg string) *Result {
	return NewResult(code, msg, nil)
}

func Unauthorized() *Result {
	return NewResult(401, "Unauthorized", nil)
}

func (r *Result) Ok() bool {
	return r.Code == 0
}
