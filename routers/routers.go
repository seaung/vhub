package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/seaung/vhub/pkg/api"
)

// 定义API版本
const (
	APIV1 = "/api/v1"
)

func InitRouters() *gin.Engine {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)

	// 静态文件和模板配置
	r.LoadHTMLGlob("templates/**/*.tmpl")
	r.Static("/static", "./static")

	// 基础路由
	r.GET("/index", api.HomePage)

	// API路由组
	v1 := r.Group(APIV1)
	{
		// RCE漏洞演示路由组
		rce := v1.Group("/rce")
		{
			rce.GET("", api.RecController)
			rce.GET("/exec", api.RceExecController)
			rce.POST("/exec", api.PostCommand)
		}

		// XSS漏洞演示路由组
		xss := v1.Group("/xss")
		{
			xss.GET("/dom", api.XSSDOMHandler)
			xss.GET("/reflected", api.XSSReflectedHandler)
			xss.GET("/stored", api.XSSStoredHandler)
			xss.POST("/stored", api.XSSStoredSubmitHandler)
		}

		// XXE漏洞演示路由组
		xxe := v1.Group("/xxe")
		{
			xxe.GET("", api.XXEHandler)
			xxe.POST("/parse", api.XXEParseHandler)
		}

		// SSRF漏洞演示路由组
		ssrf := v1.Group("/ssrf")
		{
			ssrf.GET("", api.SSRFHandler)
			ssrf.GET("/request", api.SSRFRequestHandler)
		}

		// SQL注入漏洞演示路由组
		sqli := v1.Group("/sqli")
		{
			sqli.GET("", api.SQLIHandler)
			sqli.POST("/login", api.SQLILoginHandler)
			sqli.GET("/search", api.SQLISearchHandler)
			sqli.POST("/update", api.SQLIUpdateHandler)
		}
	}

	return r
}
