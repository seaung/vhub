package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SQLControllers(c *gin.Context) {
	c.HTML(http.StatusOK, "sql.html", gin.H{
		"title": "SQL",
	})
}
