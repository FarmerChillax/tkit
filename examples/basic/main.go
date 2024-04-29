package main

import (
	"log"

	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit/app"
	"github.com/gin-gonic/gin"
)

type PingRequest struct{}

type PingResponse struct {
	Msg string `json:"message"`
}

func PingHandler(ctx *gin.Context, req *PingRequest) (*PingResponse, error) {
	tkit.Logger.Infof(ctx, "headers: %v", ctx.Request.Header)
	return &PingResponse{Msg: "pong"}, nil
}

type ErrorRequest struct{}

type ErrorResponse struct{}

func ErrorHandler(ctx *gin.Context, req *ErrorRequest) (*ErrorResponse, error) {
	tkit.Logger.Infof(ctx, "headers: %v", ctx.Request.Header)
	return nil, tkit.NewError(500, 1001, "error")
}

func main() {
	builder, err := app.NewBuilder(&tkit.Application{
		Name: "basic-demo",
	})
	if err != nil {
		log.Fatalln("app.NewBuilder err: ", err)
	}

	// 设置成功处理器
	tkit.SetSuccessHandler(func(ctx *gin.Context, data any) any {
		return tkit.ResponseDTO{
			Code:   200,
			Msg:    "success",
			Data:   data,
			Status: "ok",
		}
	})
	// 设置错误处理器
	tkit.SetErrorHandler(func(err tkit.Error) (int, any) {
		return err.StatusCode(), tkit.ResponseDTO{
			Code:   err.Code(),
			Msg:    err.Error(),
			Status: "error abort",
		}
	})

	err = builder.ListenGinServer(&tkit.GinApplication{
		RegisterHttpRoute: func(e *gin.Engine) error {
			// 注册路由，并使用 wrap 方法进行请求参数解析和返回值封装
			e.GET("/ping", tkit.Wrap(PingHandler))
			e.GET("/error", tkit.Wrap(ErrorHandler))
			return nil
		},
	})
	if err != nil {
		log.Fatalln("builder.ListenGinServer err: ", err)
	}
}
