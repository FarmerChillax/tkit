package config

import "github.com/FarmerChillax/tkit/pkg/utils"

func InitGlobalConfig(conf *Config) (*Config, error) {
	if conf != nil {
		return conf, nil
	}

	conf = getConfigFromEnv()
	return conf, nil
}

func getConfigFromEnv() *Config {
	config := &Config{
		Timeout: int64(utils.GetEnvIntByDefault("TKIT_TIMEOUT", 10)),
	}

	config.Database = getDatabaseConfigFromEnv()
	config.Redis = getRedisConfigFromEnv()
	config.Logger = getLoggerConfigFromEnv()
	config.Otel = getOtelConfigFromEnv()

	return config
}
