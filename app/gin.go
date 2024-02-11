package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit/internal/metrics"
	"github.com/FarmerChillax/tkit/internal/middlewares"
	"github.com/FarmerChillax/tkit/pkg/helper"
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

func (app *Builder) ListenGinServer(ginApp *tkit.GinApplication) error {
	ginApp.Application = app.Application
	engine := gin.New()

	if ginApp.TracerLogger == nil {
		ginApp.TracerLogger = helper.NewTracerLogger()
	}

	// 服务注册
	// todo...

	// 注册中间件
	ginMiddleware := middlewares.NewWithGin(ginApp)
	ginMiddleware.Register(engine)

	// 注册公共指标
	// 暂不将路径公开配置，后续再考虑
	engine = metrics.RegisterGin(engine)

	// 注册路由
	if ginApp.RegisterHttpRoute != nil {
		err := ginApp.RegisterHttpRoute(engine)
		if err != nil {
			return fmt.Errorf("ginApp.RegisterHttpRoute err: %v", err)
		}
	}

	// Gin 服务启动
	addr := net.JoinHostPort(ginApp.Application.Host, strconv.Itoa(int(ginApp.Application.Port)))
	server := http.Server{
		Addr:    addr,
		Handler: engine,
	}

	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		ginApp.TracerLogger.Infof("Start http server listen on: %s .", addr)
		return server.ListenAndServe()
	})

	// 服务退出处理事项
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)
	var stopSignal os.Signal
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			return nil
		case stopSignal = <-stopChan:
			// 程序退出
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()
			if err := server.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}

			return nil
		}
	})

	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	if ginApp.RegisterGracefulStopHandler != nil {
		if err := ginApp.RegisterGracefulStopHandler(stopSignal); err != nil {
			return fmt.Errorf("ginApp.RegisterGracefulStopHandler err: %v", err)
		}
	}

	return nil
}
