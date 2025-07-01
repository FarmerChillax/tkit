package tkit

import (
	"context"
	"errors"
	"fmt"

	"gitlab-ci.xiaopeng.local/voyager/logan-sdk-go/pkg/xtool"
	"golang.org/x/sync/errgroup"
)

// 应用元数据
type metadata struct {
	Name    string
	Version string
}

type Application struct {
	metadata
	// application ctx
	ctx context.Context

	StartupProcess func(ctx context.Context) error
	// 回调接口
	Callbacks  map[CallbackPoint]func(ctx context.Context) error
	AppOptions map[ApplicationOptionType]any
}

func validateApplication(app *Application) error {
	if app == nil {
		return errors.New("Application is nil")
	}

	if app.Name == "" {
		// 如果没有设置应用名称，则从环境变量中获取
		app.Name = xtool.GetEnvByDefault(EnvKeyApplicationName, "")
		if app.Name == "" {
			return errors.New("Application name is empty")
		}
	}

	// 设置应用版本信息
	if app.Version == "" {
		app.Version = xtool.GetEnvByDefault(EnvKeyApplicationVersion, "0.1.0")
	}

	if app.Callbacks == nil {
		app.Callbacks = make(map[CallbackPoint]func(context.Context) error, 0)
	}

	return nil
}

func NewApplication(app *Application) (*Application, error) {
	// create signal handler
	ctx := SetupSignalHandler()

	return NewApplicationWithCtx(ctx, app)
}

func NewApplicationWithCtx(ctx context.Context, app *Application) (*Application, error) {
	// overwrite the application ctx
	app.ctx = ctx

	// validate application
	if err := validateApplication(app); err != nil {
		panic(fmt.Errorf("validate application error: %w", err))
	}

	// configer, err := extraceConfiger(app.AppOptions)
	// if err != nil {
	// 	return nil, fmt.Errorf("extraceConfig err: %w", err)
	// }

	// if configer != nil {
	// 	Configer = configer
	// 	// 加载全局配置
	// 	err = loadGlobalConfig()
	// 	if err != nil {
	// 		return nil, fmt.Errorf("load global config error: %w", err)
	// 	}
	// }

	// 全局组件初始化
	if err := app.registerGlobalComponent(); err != nil {
		return nil, fmt.Errorf("register global component err: %w", err)
	}

	// run user default StartupProcess func
	if app.StartupProcess != nil {
		if err := app.StartupProcess(ctx); err != nil {
			return nil, fmt.Errorf("run startupProcess func err: %w", err)
		}
	}

	if err := app.endCallback(ctx); err != nil {
		return nil, fmt.Errorf("endCallback err: %w", err)
	}

	return app, nil
}

// 注册全局组件
// e.g. otel、logger、east-west Gateway、Database、Redis...
//
// note: 该版本临时服用原有方法，后续根据使用方式完善后需要 clean code，最终保留一种合理的方法
func (app *Application) registerGlobalComponent() error {
	// // register otel
	// _, err := RegisterTracer(app.Name, xtool.GetEnvByDefault(constant.EnvKeyOtelExporter, constant.DefaultOtelExporter))
	// if err != nil {
	// 	return fmt.Errorf("register otel tracer err: %w", err)
	// }

	// // register logger
	// if err := RegisterLogger(&LoggerConf); err != nil {
	// 	return fmt.Errorf("register logger err: %w", err)
	// }

	// // register database
	// if DatabaseConf != nil && DatabaseConf.Driver != "" {
	// 	if err = InitDatabase(); err != nil {
	// 		return fmt.Errorf("init database error: %v", err)
	// 	}
	// 	Logger.Infof(context.TODO(), "init database success")
	// }

	// // register redis
	// if RedisConf != nil && RedisConf.Addr != "" {
	// 	if err = InitRedis(); err != nil {
	// 		return fmt.Errorf("init redis error: %v", err)
	// 	}
	// 	Logger.Infof(context.TODO(), "init redis success")
	// }

	return nil
}

// 启动 runtime
func (app *Application) Run(runtimes ...Runtime) error {
	errg, ctx := errgroup.WithContext(app.ctx)
	for _, rt := range runtimes {
		// 注入 application 对象
		if err := rt.setApplication(app); err != nil {
			return fmt.Errorf("runtime[%s]: inject application to runtime err: %w",
				rt.name(), err)
		}

		// 启动 runtime
		errg.Go(func() error {
			err := rt.Run(ctx)
			if err != nil {
				return fmt.Errorf("runtime[%s] running err: %w", rt.name(), err)
			}
			return nil
		})
	}

	if err := errg.Wait(); err != nil {
		return fmt.Errorf("errgroup recive an err: %w", err)
	}

	return nil
}
