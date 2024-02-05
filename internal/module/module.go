package module

import (
	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit/config"
)

func Register(conf *config.Config) (err error) {
	// 初始化数据库配置
	if conf.Database != nil && conf.Database.Driver != "" && conf.Database.Dsn != "" {
		tkit.Database, err = NewDatabase(conf.Database)
		if err != nil {
			return err
		}
	}

	// 初始化 redis 配置
	if conf.Redis != nil && conf.Redis.Addr != "" {
		tkit.Redis, err = NewRedis(conf.Redis)
		if err != nil {
			return err
		}
	}

	// 初始化 HTTP 客户端配置

	return nil
}
