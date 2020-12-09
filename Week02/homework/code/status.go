package code

import "fmt"

var _ error = &StatusError{}

// StatusError 返回通用状态码
type StatusError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Error ..
func (e *StatusError) Error() string {
	return fmt.Sprintf("error: code = %d desc = %s", e.Code, e.Msg)
}

// Error 示例话错误信息
func Error(code int, msg string) error {
	return &StatusError{Code: code, Msg: msg}
}
