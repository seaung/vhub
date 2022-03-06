package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SSRFControllers(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/ssrf.html", gin.H{
		"title": "SSRF",
	})
}
