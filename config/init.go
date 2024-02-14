package config

import "github.com/FarmerChillax/tkit/pkg/utils"

func InitGlobalConfig(conf *Config) (*Config, error) {
	if conf != nil {
		config = conf
		return config, nil
	}

	config = getConfigFromEnv()
	return config, nil
}

func getConfigFromEnv() *Config {
	config = &Config{
		Timeout: int64(utils.GetEnvIntByDefault("TKIT_TIMEOUT", 10)),
	}

	config.Database = getDatabaseConfigFromEnv()
	config.Redis = getRedisConfigFromEnv()
	config.Logger = getLoggerConfigFromEnv()
	config.Otel = getOtelConfigFromEnv()

	return config
}
