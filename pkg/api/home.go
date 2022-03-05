package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home/index.html", gin.H{"title": "VHub a web broken application"})
}
