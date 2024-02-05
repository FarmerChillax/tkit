package middlewares

import (
	"context"
	"time"

	"github.com/FarmerChillax/tkit"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type GinMiddleware struct {
	*tkit.GinApplication
}

func NewWithGin(ginApp *tkit.GinApplication) *GinMiddleware {
	return &GinMiddleware{
		GinApplication: ginApp,
	}
}

func (m *GinMiddleware) Register(engine *gin.Engine) {
	// 注册公共中间件
	engine.Use(gin.Recovery())
	engine.Use(otelgin.Middleware(m.Application.Name))
	engine.Use(m.AccessLog())

	if m.Application.Config != nil {
		engine.Use(ContextTimeout(time.Second * time.Duration(m.Application.Config.Timeout)))
	}
}

func (m *GinMiddleware) AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		startT := time.Now()
		c.Next()
		latency := time.Since(startT)
		logInfo := "access request, method: %v, uri: %v, status_code: %v, latency: %v"
		m.TracerLogger.WithContext(c.Request.Context()).Infof(logInfo, c.Request.Method, c.Request.RequestURI, c.Writer.Status(), latency)
	}
}

func ContextTimeout(timeout time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), timeout)
		defer cancel()
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
