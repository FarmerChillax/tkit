package config

import (
	"github.com/sirupsen/logrus"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var config *Config

type Config struct {
	// 全局请求超时
	Timeout int64
	// RequestTimeout int64
	// ReadTimeout    int64
	// WriteTimeout   int64
	Database *DatabseConfig `json:"mysql,omitempty" mapstructure:"database"`
	Redis    *RedisConfig   `json:"redis,omitempty"`
	Logger   *LoggerConfig  `json:"logger,omitempty"`
	Otel     *OtelConfig    `json:"otel,omitempty"`
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

func (c *Config) GetDatabase() *DatabseConfig {
	if c.Database == nil {
		c.Database = &DatabseConfig{
			Driver: "sqlite3",
			Dsn:    ":memory:",
		}
	}
	return c.Database
}

type RedisConfig struct {
	Addr     string
	Password string
	DB       int
	PoolSize int
	MaxIdle  int
}

func (c *Config) GetRedis() *RedisConfig {
	if c.Redis == nil {
		c.Redis = &RedisConfig{}
	}
	return c.Redis
}

type LoggerConfig struct {
	Formatter logrus.Formatter
	RootPath  string
}

func (c *Config) GetLoggerConfig() *LoggerConfig {
	if c.Logger == nil {
		c.Logger = &LoggerConfig{}
	}
	return c.Logger
}

type OtelConfig struct {
	Exporter sdktrace.SpanExporter
}

func (c *Config) GetOtelConfig() *OtelConfig {
	if c.Otel == nil {
		c.Otel = &OtelConfig{}
	}
	return c.Otel
}
