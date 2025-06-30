package tkit

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error interface {
	error
	// WithRequestID 设置当前请求的唯一ID
	// WithRequestID(requestId string) Error
	// 设置 http status code
	WithStatusCode(statusCode int) Error
	// 设置错误描述
	WithMsg(msg string) Error
	WithMsgf(format string, args ...interface{}) Error

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

// 从字节流中解析错误，并返回是否为 tkit 的错误类型
func NewErrorFromBytes(data []byte) (Error, bool) {
	e := &err{
		statusCode: http.StatusInternalServerError,
		Code:       0,
		Msg:        "未知错误",
	}

	err := json.Unmarshal(data, &e)
	if err != nil {
		return e.WithError(err).WithMsg(err.Error()), false
	}

	return e, true
}

// 从字符串中解析错误，并返回是否为 tkit 的错误类型
func NewErrorFromString(data string) (Error, bool) {
	return NewErrorFromBytes([]byte(data))
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

func (e *err) WithMsgf(format string, args ...interface{}) Error {
	newErr := e.clone()
	newErr.Msg = fmt.Sprintf(format, args...)
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
