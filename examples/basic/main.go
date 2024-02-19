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

func main() {
	builder, err := app.NewBuilder(&tkit.Application{
		Name: "basic-demo",
	})
	if err != nil {
		log.Fatalln("app.NewBuilder err: ", err)
	}

	err = builder.ListenGinServer(&tkit.GinApplication{
		RegisterHttpRoute: func(e *gin.Engine) error {
			// 注册路由，并使用 wrap 方法进行请求参数解析和返回值封装
			e.GET("/ping", tkit.Wrap(PingHandler))
			return nil
		},
	})
	if err != nil {
		log.Fatalln("builder.ListenGinServer err: ", err)
	}
}
