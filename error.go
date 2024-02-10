package tkit

import "net/http"

type Error interface {
	error
	// WithID 设置当前请求的唯一ID
	WithID(id string) Error
	// 设置 http status code
	WithStatusCode(statusCode int) Error
	// 设置错误描述
	WithMsg(msg string) Error

	// 返回业务编码
	Code() int
	// 返回对应的 http 状态码
	StatusCode() int
}

type err struct {
	statusCode int
	code       int    // 业务编码
	msg        string // 错误描述
	id         string // 当前请求的唯一ID
}

// NewError 返回一个新的自定义错误
// NewError 返回一个新的自定义错误
func NewError(httpStatusCode, businessCode int, msg string) Error {
	return &err{
		statusCode: httpStatusCode,
		code:       businessCode,
		msg:        msg,
	}
}

func (e *err) Error() string {
	return e.msg
}

func (e *err) Code() int {
	return e.code
}

func (e *err) StatusCode() int {
	if e.statusCode > 0 {
		return e.statusCode
	}

	return http.StatusInternalServerError
}

func (e *err) WithID(id string) Error {
	e.id = id
	return e
}

func (e *err) WithStatusCode(statusCode int) Error {
	e.statusCode = statusCode
	return e
}

func (e *err) WithMsg(msg string) Error {
	e.msg = msg
	return e
}
