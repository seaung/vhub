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
	r.GET("/xss", api.XSSCotrollers)
	r.GET("/xss-demon-01", api.XSS01Controllers)
	r.GET("/xxe", api.XXEControllers)
	r.GET("/ssrf", api.SSRFControllers)
	r.GET("/sql", api.SQLControllers)
	r.GET("/fileinclude", api.FileIncludeController)
	r.GET("/csrf", api.CSRFControllers)

	return r
}
