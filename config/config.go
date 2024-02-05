package config

import "github.com/sirupsen/logrus"

type Config struct {
	// 全局请求超时
	Timeout int64
	// RequestTimeout int64
	// ReadTimeout    int64
	// WriteTimeout   int64
	Database *DatabseConfig `json:"mysql,omitempty" mapstructure:"database"`
	Redis    *RedisConfig   `json:"redis,omitempty"`
	Logger   *LoggerConfig  `json:"logger,omitempty"`
}

var config *Config

func Get() *Config {
	return config
}

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

func GetDatabase() *DatabseConfig {
	if config.Database == nil {
		config.Database = &DatabseConfig{
			Driver: "sqlite3",
		}
	}
	return config.Database
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
	MaxIdle  int
}

func GetRedis() *RedisConfig {
	// if config.Redis == nil {
	// 	return &defaultRedisConfig
	// }
	return config.Redis
}

type LoggerConfig struct {
	Formatter logrus.Formatter
	RootPath  string
}

func GetLoggerConfig() *LoggerConfig {
	if config.Logger == nil {
		config.Logger = &LoggerConfig{}
	}
	return config.Logger
}
