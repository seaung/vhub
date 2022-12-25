package api

import (
	"net/http"

	"os/exec"

	"github.com/gin-gonic/gin"
)

var cmdList = []VulnsDesc{
	{Name: "RCE", Desc: "RCE漏洞，可以让攻击者直接向后台服务器远程注入操作系统命令或者代码，从而控制后台系统.", VulnLinks: "/rce"},
	{Name: "System", Desc: "RCE漏洞，可以让攻击者直接向后台服务器远程注入操作系统命令或者代码，从而控制后台系统.", VulnLinks: "/system"},
}

func CmdControllers(c *gin.Context) {
	c.HTML(http.StatusOK, "cmd.html", gin.H{
		"title": "Cmd",
		"vulns": cmdList,
	})
}

func RCEVulns(c *gin.Context) {
	command := c.DefaultQuery("cmd", "echo")

	args := c.DefaultQuery("", "")

	exec.Command(command, args)

	c.HTML(http.StatusOK, "rce.html", gin.H{})
}
