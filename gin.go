package tkit

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type GinApplication struct {
	Application                 *Application
	TracerLogger                *logrus.Logger
	RegisterHttpRoute           func(*gin.Engine) error
	RegisterGracefulStopHandler func(sig ...os.Signal) error
}
