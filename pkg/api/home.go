package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"pages/index.tmpl",
		vulns,
	)
}
