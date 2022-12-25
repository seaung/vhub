package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func FileIncludeController(c *gin.Context) {
	c.HTML(http.StatusOK, "fileinclude.html", gin.H{
		"title": "file include",
	})
}
