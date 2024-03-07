package config

import (
	"github.com/FarmerChillax/tkit/pkg/utils"
	"github.com/sirupsen/logrus"
)

type LoggerConfig struct {
	Formatter logrus.Formatter
	RootPath  string
	// OutputPath   string
	ReportCaller bool
	Level        uint32
}

func getLoggerConfigFromEnv() *LoggerConfig {
	return &LoggerConfig{
		RootPath: utils.GetEnvByDefault("TKIT_LOGGER_ROOT_PATH", "./logs"),
		// OutputPath:   utils.GetEnvByDefault("TKIT_LOGGER_OUTPUT_PATH", "./logger.log"),
		ReportCaller: utils.GetEnvBoolByDefault("TKIT_LOGGER_REPORT_CALLER", false),
		Level:        uint32(utils.GetEnvIntByDefault("TKIT_LOGGER_LEVEL", int(logrus.InfoLevel))),
	}
}
