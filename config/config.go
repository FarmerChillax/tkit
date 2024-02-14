package config

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

// 获取数据库配置
// 如果没有配置，则从环境变量读取
// 如果环境变量没有配置，则使用默认配置
func (c *Config) GetDatabase() *DatabseConfig {
	if c.Database == nil {
		c.Database = getDatabaseConfigFromEnv()
	}
	return c.Database
}

// 获取 redis 配置
// 如果没有配置，则从环境变量读取
// 如果环境变量没有配置，则使用默认配置
func (c *Config) GetRedis() *RedisConfig {
	if c.Redis == nil {
		c.Redis = getRedisConfigFromEnv()
	}
	return c.Redis
}

// 获取 logger 配置
// 如果没有配置，则从环境变量读取
// 如果环境变量没有配置，则使用默认配置
func (c *Config) GetLoggerConfig() *LoggerConfig {
	if c.Logger == nil {
		c.Logger = getLoggerConfigFromEnv()
	}
	return c.Logger
}

// 获取 otel 配置
// 如果没有配置，则从环境变量读取
// 如果环境变量没有配置，则使用默认配置
func (c *Config) GetOtelConfig() *OtelConfig {
	if c.Otel == nil {
		c.Otel = getOtelConfigFromEnv()
	}
	return c.Otel
}
