package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"pages/index.tmpl",
		gin.H{
			"title": "漏洞演示平台",
			"vulns": vulns,
			"description": "一个用于演示常见Web安全漏洞的平台",
			"version": "1.0.0",
		},
	)
}
