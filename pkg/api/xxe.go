package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func XXEControllers(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/xxe.html", gin.H{
		"title": "XXE",
	})
}
