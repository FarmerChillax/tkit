package tkit

import (
	"github.com/gin-gonic/gin"
)

// type ResponseIface interface {
// 	// 返回响应结果
// 	Result() any
// 	// 响应器名称
// 	Name() string
// }

// var Response ResponseIface

type ResponseDataIface interface {
	GetStatus() string
}

type ResponseDTO struct {
	Code   int    `json:"code"`
	Msg    string `json:"msg"`
	Data   any    `json:"data"`
	Status string `json:"status"`
}

func (r *ResponseDTO) GetStatus() string {
	if r.Status != "" {
		return r.Status
	}
	return "ok"
}

// Result default func
// Result 方法的默认值包装方法
func ResultData(ctx *gin.Context, data interface{}) {
	Result(ctx, defautlHttpStatusCode, defautlBizStatusCode, data, "", "success")
}

// Result result template
func Result(ctx *gin.Context, httpCode, code int, data interface{}, message string, status string) {
	response := ResponseDTO{
		Code:   code,
		Msg:    message,
		Data:   data,
		Status: status,
	}
	ctx.JSON(httpCode, response)
}

// ResultError result failed
func ResultError(ctx *gin.Context, err Error) {
	response := ResponseDTO{
		Code:   err.Code(),
		Msg:    err.Error(),
		Data:   nil,
		Status: "failed",
	}
	ctx.AbortWithStatusJSON(err.StatusCode(), response)
}

// 通用 json 响应
func CommonJsonResult(ctx *gin.Context, code int, resp any) {
	ctx.JSON(code, resp)
}
