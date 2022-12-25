package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecController(c *gin.Context) {
	c.HTML(http.StatusOK, "pages/rce.html", gin.H{
		"title": "rec page",
	})
}
