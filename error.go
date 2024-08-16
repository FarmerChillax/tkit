package tkit

import "net/http"

type Error interface {
	error
	// WithRequestID 设置当前请求的唯一ID
	// WithRequestID(requestId string) Error
	// 设置 http status code
	WithStatusCode(statusCode int) Error
	// 设置错误描述
	WithMsg(msg string) Error
	// 设置错误栈
	WithError(err error) Error

	// 返回业务编码
	GetCode() int
	// 返回对应的 http 状态码
	GetStatusCode() int
	// 返回错误描述
	GetMsg() string

	// 返回原始错误
	Unwrap() error
}

type err struct {
	statusCode int
	Code       int    `json:"code,omitempty"` // 业务编码
	Msg        string `json:"msg,omitempty"`  // 错误描述
	// requestId  string // 当前请求的唯一ID
	err error
}

// NewError 返回一个新的自定义错误
func NewError(httpStatusCode, businessCode int, msg string) Error {
	return &err{
		statusCode: httpStatusCode,
		Code:       businessCode,
		Msg:        msg,
	}
}

func (e *err) Error() string {
	return e.Msg
}

func (e *err) GetCode() int {
	return e.Code
}

func (e *err) GetStatusCode() int {
	if e.statusCode > 0 {
		return e.statusCode
	}

	return http.StatusInternalServerError
}

func (e *err) GetMsg() string {
	return e.Msg
}

func (e *err) clone() *err {
	return &err{
		statusCode: e.statusCode,
		Code:       e.Code,
		Msg:        e.Msg,
		// requestId:  e.requestId,
	}
}

// func (e *err) WithRequestID(requestId string) Error {
// 	newErr := e.clone()
// 	newErr.requestId = requestId
// 	return newErr
// }

func (e *err) WithStatusCode(statusCode int) Error {
	newErr := e.clone()
	newErr.statusCode = statusCode
	return newErr
}

func (e *err) WithMsg(msg string) Error {
	newErr := e.clone()
	newErr.Msg = msg
	return newErr
}
func (e *err) WithError(err error) Error {
	newErr := e.clone()
	newErr.err = err
	return newErr
}

func (e *err) Unwrap() error {
	return e.err
}
