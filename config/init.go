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

	return config
}

func getDatabaseConfigFromEnv() *DatabseConfig {
	dbConfig := DatabseConfig{
		Dsn:               utils.GetEnvByDefault("TKIT_DATABASE_DSN", ""),
		Driver:            utils.GetEnvByDefault("TKIT_DATABASE_DRIVER", ""),
		Loc:               utils.GetEnvByDefault("TKIT_DATABASE_LOC", "Local"),
		ParseTime:         utils.GetEnvByDefault("TKIT_DATABASE_PARSE_TIME", "true"),
		Timeout:           int64(utils.GetEnvIntByDefault("TKIT_DATABASE_TIMEOUT", 10)),
		MaxOpen:           utils.GetEnvIntByDefault("TKIT_DATABASE_MAX_OPEN", 30),
		MaxIdle:           utils.GetEnvIntByDefault("TKIT_DATABASE_MAX_IDLE", 10),
		ConnMaxLifeSecond: utils.GetEnvIntByDefault("TKIT_DATABASE_CONN_MAX_LIFE_SECOND", 60),
	}
	return &dbConfig
}

func getRedisConfigFromEnv() *RedisConfig {
	redisConf := RedisConfig{
		Addr:     utils.GetEnvByDefault("TKIT_REDIS_ADDR", "127.0.0.1"),
		Password: utils.GetEnvByDefault("TKIT_REDIS_PASSWORD", ""),
		DB:       utils.GetEnvIntByDefault("TKIT_REDIS_DB", 0),
		PoolSize: utils.GetEnvIntByDefault("TKIT_REDIS_POOL_SIZE", 30),
		MaxIdle:  utils.GetEnvIntByDefault("TKIT_REDIS_MAX_IDLE", 10),
	}
	return &redisConf
}

func getLoggerConfigFromEnv() *LoggerConfig {
	return &LoggerConfig{
		RootPath: utils.GetEnvByDefault("TKIT_LOGGER_ROOT_PATH", "./logs"),
	}
}
