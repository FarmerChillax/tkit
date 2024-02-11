package app

import (
	"errors"
	"log"

	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit/config"
	"github.com/FarmerChillax/tkit/internal/module"
	"github.com/FarmerChillax/tkit/internal/otel"
	"github.com/FarmerChillax/tkit/internal/xlog"
)

type Builder struct {
	Application *tkit.Application
}

func NewBuilder(app *tkit.Application) (*Builder, error) {
	// tkit.ApplicationInstance = app
	if err := validateAndMergeAppConfig(app); err != nil {
		return nil, err
	}

	// 初始化 config
	conf, err := config.InitGlobalConfig(app.Config)
	if err != nil {
		log.Println("InitGlobalConfig err: ", err)
		return nil, err
	}
	err = runCallback(tkit.POSITION_GLOBAL_CONFIG, app.RegisterCallback)
	if err != nil {
		log.Println("runCallback POSITION_GLOBAL_CONFIG err: ", err)
		return nil, err
	}

	// 初始化 otel
	_, err = otel.RegisterTracer(app.Name, conf.GetOtelConfig().Exporter)
	if err != nil {
		log.Println("RegisterTracer err: ", err)
		return nil, err
	}

	// 初始化日志
	err = xlog.Register(conf.GetLoggerConfig())
	if err != nil {
		log.Println("Register Logger err: ", err)
		return nil, err
	}
	err = runCallback(tkit.POSITION_INIT_LOGGER, app.RegisterCallback)
	if err != nil {
		log.Println("runCallback POSITION_INIT_LOGGER err: ", err)
		return nil, err
	}

	// 初始化内置组件
	if err := module.Register(conf); err != nil {
		log.Println("Register Modules err: ", err)
		return nil, err
	}
	err = runCallback(tkit.POSITION_MODULE_REGISTER, app.RegisterCallback)
	if err != nil {
		log.Println("runCallback POSITION_MODULE_REGISTER err: ", err)
		return nil, err
	}

	// 加载用户配置
	if app.LoadConfig != nil {
		err := app.LoadConfig()
		if err != nil {
			return nil, err
		}
		// 执行加载配置回调操作
		if err := runCallback(tkit.POSITION_LOAD_CONFIG, app.RegisterCallback); err != nil {
			return nil, err
		}
	}

	// 设置常量
	if app.SetupVars != nil {
		if err := app.SetupVars(); err != nil {
			return nil, err
		}
		// 执行设置常量的回调操作
		if err := runCallback(tkit.POSITION_SETUP_VARS, app.RegisterCallback); err != nil {
			return nil, err
		}
	}

	if err := runCallback(tkit.POSITION_NEW, app.RegisterCallback); err != nil {
		return nil, err
	}

	return &Builder{
		Application: app,
	}, nil
}

func validateAndMergeAppConfig(app *tkit.Application) error {
	if app.Name == "" {
		return errors.New("application name can't not be empty")
	}

	// flag.Parse()
	// todo... 从命令行参数中获取配置

	if app.Port > 65535 || app.Port <= 0 {
		log.Println("port is invalid, use default port 6000")
		app.Port = 6000
	}

	if app.Host == "" {
		app.Host = "127.0.0.1"
	}

	return nil
}

func runCallback(position tkit.CallbackPosition, callbacks map[tkit.CallbackPosition]tkit.CallbackFunc) error {
	if f, ok := callbacks[position]; ok {
		return f()
	}

	return nil
}
