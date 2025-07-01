package tkit

import (
	"context"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type CallbackPosition int

const (
	// 在 InitGlobalConfig 后调用
	POSITION_GLOBAL_CONFIG CallbackPosition = iota + 1
	// 在 InitGlobalLogger 后调用
	POSITION_INIT_LOGGER
	// 在 Module Register 后调用
	POSITION_MODULE_REGISTER
	// 调用 LoadConfig 方法后
	POSITION_LOAD_CONFIG
	// 调用 SetupVars 方法后
	POSITION_SETUP_VARS
	// 调用 New 方法后
	POSITION_NEW
)

type CallbackFunc func() error

type Configer interface{}

type DatabaseIface interface {
	Get(ctx context.Context) *gorm.DB
}

var Database DatabaseIface

type RedisConn interface {
	Get(ctx context.Context) *redis.Client
}

var Redis RedisConn

type LoggerIface interface {
	// 获取日志实例
	// Get(ctx context.Context) LoggerIface

	// debug 日志
	// Debug(ctx context.Context, args ...interface{})
	// Debugf(ctx context.Context, format string, args ...interface{})

	// 普通日志
	Info(ctx context.Context, args ...interface{})
	Infof(ctx context.Context, format string, args ...interface{})

	// 警告日志
	Warn(ctx context.Context, args ...interface{})
	Warnf(ctx context.Context, format string, args ...interface{})

	// 错误日志
	Error(ctx context.Context, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
}

var Logger LoggerIface
