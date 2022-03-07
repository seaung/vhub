package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func XSSCotrollers(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/xss.html", gin.H{
		"title": "xss",
	})
}

func XSS01Controllers(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/xss-01.html", gin.H{
		"title": "xss 01 demo",
	})
}
