package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var cmdList = []VulnsDesc{
	{Name: "RCE", Desc: "RCE漏洞，可以让攻击者直接向后台服务器远程注入操作系统命令或者代码，从而控制后台系统.", VulnLinks: "/rce"},
	{Name: "System", Desc: "RCE漏洞，可以让攻击者直接向后台服务器远程注入操作系统命令或者代码，从而控制后台系统.", VulnLinks: "/system"},
}

func CmdControllers(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/cmd.html", gin.H{
		"title": "Cmd",
		"vulns": cmdList,
	})
}
