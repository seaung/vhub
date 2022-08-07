package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var xssLists = []VulnsDesc{
	{Name: "xss1", Desc: "xss", VulnLinks: "/xss1"},
	{Name: "xss2", Desc: "xss", VulnLinks: "/xss2"},
	{Name: "xss3", Desc: "xss", VulnLinks: "/xss3"},
	{Name: "xss4", Desc: "xss", VulnLinks: "/xss4"},
}

func XSSCotrollers(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/xss.html", gin.H{
		"title": "xss",
		"vulns": xssLists,
	})
}

func XSS01Controllers(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/xss-01.html", gin.H{
		"title": "xss 01 demo",
	})
}
