package helper

import (
	"fmt"
	"strings"
	"time"

	"github.com/FarmerChillax/tkit/config"
	"github.com/FarmerChillax/tkit/pkg/utils"
	"github.com/glebarez/sqlite"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewGormDB(conf *config.DatabseConfig) (*gorm.DB, error) {
	dial, err := getDialByDriver(conf.Driver, conf.Dsn)
	if err != nil {
		return nil, err
	}

	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if utils.IsDev() {
		db.Logger = db.Logger.LogMode(logger.Info)
	} else {
		db.Logger = db.Logger.LogMode(logger.Error)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	var maxIdle, maxOpen int = 10, 30
	if conf.MaxOpen > 0 && conf.MaxIdle > 0 {
		maxIdle = conf.MaxIdle
		maxOpen = conf.MaxOpen
	}
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)

	if conf.ConnMaxLifeSecond > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(conf.ConnMaxLifeSecond) * time.Second)
	}

	return db, nil
}

func getDialByDriver(driver, dsn string) (dial gorm.Dialector, err error) {
	// add more driver here
	// for example postgresql, sqlserver, oracle, etc.
	switch strings.ToLower(driver) {
	case "mysql":
		return mysql.Open(dsn), nil
	case "sqlite3":
		return sqlite.Open(dsn), nil
	}
	return nil, fmt.Errorf("not support driver: %s", driver)
}
