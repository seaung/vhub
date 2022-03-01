package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/seaung/vhub/pkg/api"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Logger(), gin.Recovery())

	r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/static", "./static")

	r.GET("/index", api.HomePage)

	return r
}
