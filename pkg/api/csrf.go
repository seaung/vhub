package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CSRFControllers(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/csrf.html", gin.H{
		"title": "CSRF",
	})
}
