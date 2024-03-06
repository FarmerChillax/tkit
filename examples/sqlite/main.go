package main

import (
	"context"
	"fmt"
	"log"

	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit/app"
	"github.com/FarmerChillax/tkit/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func (u User) TableName() string {
	return "user"
}

func main() {
	appBuilder, err := app.NewBuilder(&tkit.Application{
		Name: "app-demo",
		Host: "0.0.0.0",
		Port: 6000,
		Config: &config.Config{
			Timeout: 120,
			Database: &config.DatabseConfig{
				Driver: "sqlite3",
				Dsn:    ":memory:",
			},
		},
		RegisterCallback: map[tkit.CallbackPosition]tkit.CallbackFunc{
			tkit.POSITION_NEW: func() error {
				// 初始化数据库
				tkit.Database.Get(context.Background()).AutoMigrate(&User{})
				return nil
			},
		},
	})
	if err != nil {
		log.Fatalln("app.New err: ", err)
	}

	err = appBuilder.ListenGinServer(&tkit.GinApplication{
		RegisterHttpRoute: func(e *gin.Engine) error {
			e.POST("/user", func(ctx *gin.Context) {
				user := User{
					Name: "tkit",
					Age:  18,
				}
				db := tkit.Database.Get(ctx.Request.Context())
				err := db.Table(User{}.TableName()).Create(&user).Error
				if err != nil {
					tkit.Logger.Errorf(ctx.Request.Context(), "db.Table.Create err: %v", err)
					ctx.JSON(500, gin.H{
						"message": err.Error(),
					})
					return
				}
				ctx.JSON(200, gin.H{
					"message": "success",
					"user":    user,
				})
			})
			e.GET("/user", func(ctx *gin.Context) {
				user := User{}
				db := tkit.Database.Get(ctx.Request.Context())
				err := db.Table(User{}.TableName()).First(&user).Error
				if err != nil {
					tkit.Logger.Errorf(ctx.Request.Context(), "db.Table.First err: %v", err)
					ctx.JSON(500, gin.H{
						"message": err.Error(),
					})
					return
				}
				ctx.JSON(200, user)
			})

			return nil
		},
	})
	if err != nil {
		log.Fatalln("appBuilder.ListenGinServer err: ", err)
	}

	fmt.Printf("%+v\n", appBuilder)
}
