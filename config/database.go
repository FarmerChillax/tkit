package config

import "github.com/FarmerChillax/tkit/pkg/utils"

type DatabseConfig struct {
	Dsn               string `json:"dsn,omitempty"`
	Driver            string `json:"driver,omitempty"`
	Loc               string `json:"loc,omitempty"`
	ParseTime         string `json:"parse_time,omitempty"`
	Timeout           int64  `json:"timeout,omitempty"`
	MaxOpen           int    `json:"max_open,omitempty"`
	MaxIdle           int    `json:"max_idle,omitempty"`
	ConnMaxLifeSecond int    `json:"conn_max_life_second,omitempty"`
}

func getDatabaseConfigFromEnv() *DatabseConfig {
	dbConfig := DatabseConfig{
		Dsn:               utils.GetEnvByDefault("TKIT_DATABASE_DSN", ":memory:"),
		Driver:            utils.GetEnvByDefault("TKIT_DATABASE_DRIVER", "sqlite3"),
		Loc:               utils.GetEnvByDefault("TKIT_DATABASE_LOC", "Local"),
		ParseTime:         utils.GetEnvByDefault("TKIT_DATABASE_PARSE_TIME", "true"),
		Timeout:           int64(utils.GetEnvIntByDefault("TKIT_DATABASE_TIMEOUT", 10)),
		MaxOpen:           utils.GetEnvIntByDefault("TKIT_DATABASE_MAX_OPEN", 1),
		MaxIdle:           utils.GetEnvIntByDefault("TKIT_DATABASE_MAX_IDLE", 1),
		ConnMaxLifeSecond: utils.GetEnvIntByDefault("TKIT_DATABASE_CONN_MAX_LIFE_SECOND", 60),
	}
	return &dbConfig
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
	MaxIdle  int
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
