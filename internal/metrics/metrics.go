package metrics

import (
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterGin(engine *gin.Engine) *gin.Engine {
	// support url path config
	// todo...

	pprof.Register(engine, "/debug/pprof")
	engine.GET("/metrics", gin.WrapH(promhttp.Handler()))
	engine.GET("/health", HealthHandler())
	return engine
}
