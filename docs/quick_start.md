# Quick Start

## 安装包

```bash
$ go get github.com/FarmerChillax/tkit
```

## 快速开始

1. 初始化应用构建器
```go
func main() {
	builder, err := app.NewBuilder(&tkit.Application{
		Name: "basic-demo", // 应用名称不能为空
	})
	if err != nil {
		log.Fatalln("app.NewBuilder err: ", err)
	}
}
```

2. 实现业务逻辑并注册 http 路由
```go
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
```

3. 运行程序
```bash
$ go run main.go
>>>
2024/02/14 14:15:16 port is invalid, use default port 6000
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /debug/pprof/             --> github.com/gin-contrib/pprof.RouteRegister.WrapF.func1 (5 handlers)
[GIN-debug] GET    /debug/pprof/cmdline      --> github.com/gin-contrib/pprof.RouteRegister.WrapF.func2 (5 handlers)
[GIN-debug] GET    /debug/pprof/profile      --> github.com/gin-contrib/pprof.RouteRegister.WrapF.func3 (5 handlers)
[GIN-debug] POST   /debug/pprof/symbol       --> github.com/gin-contrib/pprof.RouteRegister.WrapF.func4 (5 handlers)
[GIN-debug] GET    /debug/pprof/symbol       --> github.com/gin-contrib/pprof.RouteRegister.WrapF.func5 (5 handlers)
[GIN-debug] GET    /debug/pprof/trace        --> github.com/gin-contrib/pprof.RouteRegister.WrapF.func6 (5 handlers)
[GIN-debug] GET    /debug/pprof/allocs       --> github.com/gin-contrib/pprof.RouteRegister.WrapH.func7 (5 handlers)
[GIN-debug] GET    /debug/pprof/block        --> github.com/gin-contrib/pprof.RouteRegister.WrapH.func8 (5 handlers)
[GIN-debug] GET    /debug/pprof/goroutine    --> github.com/gin-contrib/pprof.RouteRegister.WrapH.func9 (5 handlers)
[GIN-debug] GET    /debug/pprof/heap         --> github.com/gin-contrib/pprof.RouteRegister.WrapH.func10 (5 handlers)
[GIN-debug] GET    /debug/pprof/mutex        --> github.com/gin-contrib/pprof.RouteRegister.WrapH.func11 (5 handlers)
[GIN-debug] GET    /debug/pprof/threadcreate --> github.com/gin-contrib/pprof.RouteRegister.WrapH.func12 (5 handlers)
[GIN-debug] GET    /metrics                  --> github.com/FarmerChillax/tkit/internal/metrics.RegisterGin.WrapH.func1 (5 handlers)
[GIN-debug] GET    /health                   --> github.com/FarmerChillax/tkit/internal/metrics.RegisterGin.HealthHandler.func2 (5 handlers)
[GIN-debug] GET    /ping                     --> main.main.func1.1 (5 handlers)
  _  __  _   _             _                    _
 | |/ / (_) | |_     ___  | |_    __ _   _ __  | |_
 | ' /  | | | __|   / __| | __|  / _` | | '__| | __|
 | . \  | | | |_    \__ \ | |_  | (_| | | |    | |_
 |_|\_\ |_|  \__|   |___/  \__|  \__,_| |_|     \__|
Formatter.Format:  map[]
{"level":"info","msg":"http server listen on: 127.0.0.1:6000 .","span_id":"0000000000000000","time":"2024-02-14T14:15:16+08:00","trace_id":"00000000000000000000000000000000"}
```

4. 请求上面定义的 `ping` 接口或者框架内部接口
```bash
curl http://localhost:6000/ping
>>> {"message":"pong"}
```

