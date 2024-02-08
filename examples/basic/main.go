package main

import (
	"log"

	"github.com/FarmerChillax/tkit"
	"github.com/FarmerChillax/tkit/app"
	"github.com/gin-gonic/gin"
)

func main() {
	builder, err := app.NewBuilder(&tkit.Application{
		Name: "basic-demo",
	})
	if err != nil {
		log.Fatalln("app.New err: ", err)
	}
	err = builder.ListenGinServer(&tkit.GinApplication{
		RegisterHttpRoute: func(e *gin.Engine) error {
			e.GET("/ping", func(c *gin.Context) {
				tkit.Logger.Infof(c.Request.Context(), "headers: %v", c.Request.Header)
				c.JSON(200, gin.H{
					"message": "pong",
				})
			})
			return nil
		},
	})
	if err != nil {
		log.Fatalln("appBuilder.ListenGinServer err: ", err)
	}

}
