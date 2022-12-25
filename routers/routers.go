package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/seaung/vhub/pkg/api"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(gin.Logger(), gin.Recovery())

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.GET("/index", api.HomePage)

	xssRoute := r.Group("/xss")
	{
		xssRoute.GET("/xss-home", api.CSRFControllers)
		xssRoute.GET("/xss-1", api.XSS01Controllers)
	}

	xxeRoute := r.Group("/xxe")
	{
		xxeRoute.GET("/xxe-home", api.XXEControllers)
	}

	r.GET("/ssrf", api.SSRFControllers)
	r.GET("/sql", api.SQLControllers)
	r.GET("/fileinclude", api.FileIncludeController)
	r.GET("/csrf", api.CSRFControllers)
	r.GET("cmd", api.CmdControllers)

	return r
}
