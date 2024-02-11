package module

import (
	"context"

	"github.com/FarmerChillax/tkit/config"
	"github.com/FarmerChillax/tkit/pkg/helper"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

// var mysqlOnce sync.Once

type databseConn struct {
	client *gorm.DB
}

func (mc *databseConn) Get(ctx context.Context) *gorm.DB {
	return mc.client.WithContext(ctx)
}

func NewDatabase(conf *config.DatabseConfig) (*databseConn, error) {
	db, err := helper.NewGormDB(conf)
	if err != nil {
		return nil, err
	}

	err = db.Use(tracing.NewPlugin())
	if err != nil {
		return nil, err
	}

	return &databseConn{
		client: db,
	}, err
}
