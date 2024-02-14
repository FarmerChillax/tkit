package config

import (
	"github.com/FarmerChillax/tkit/pkg/utils"
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Formatter logrus.Formatter
	RootPath  string
}

func getLoggerConfigFromEnv() *LoggerConfig {
	return &LoggerConfig{
		RootPath: utils.GetEnvByDefault("TKIT_LOGGER_ROOT_PATH", "./logs"),
	}
}
