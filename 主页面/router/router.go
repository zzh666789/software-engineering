package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.GET("/", controller.Index1)

	e.GET("/register", controller.GoRegister)
	e.POST("/register", controller.Register)

	e.GET("/login", controller.GoLogin)
	e.POST("/login", controller.Login)

	e.GET("/index", controller.GoIndex)
	e.POST("/index", controller.GoIndex)

	e.POST("/level1", controller.Golevel1)
	e.POST("/Level1answer", controller.Golevel1Answer)

	e.POST("/level2", controller.Golevel2)
	e.POST("/Level2answer", controller.Golevel2Answer)

	e.POST("/level3", controller.Golevel3)
	e.POST("/Level3answer", controller.Golevel3Answer)
	e.Run()
}
