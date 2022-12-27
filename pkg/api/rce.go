package api

import (
	"net/http"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func RecController(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/rce.tmpl", gin.H{
		"title": "rec page",
	})
}

func RceExecController(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/rec_exec.tmpl", gin.H{
		"title": "可执行命令",
	})
}

func PostCommand(c *gin.Context) {
	cmd := c.PostForm("cmd")

	out := exec.Command("bash", cmd)

	err := out.Start()
	if err != nil {
		return
	}

	c.HTML(http.StatusOK, "vulns/rce_exec.tmpl", gin.H{
		"res": out,
	})
}
