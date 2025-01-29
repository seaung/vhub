package api

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User 定义用于演示的用户结构
type User struct {
	XMLName xml.Name `xml:"user"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}

// XXEHandler 处理XXE漏洞演示页面
func XXEHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/xxe.tmpl", gin.H{
		"title": "XXE漏洞演示",
	})
}

// XXEParseHandler 处理XML数据解析（不安全的方式）
func XXEParseHandler(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.HTML(http.StatusOK, "vulns/xxe.tmpl", gin.H{
			"title": "XXE漏洞演示",
			"error": "读取请求数据失败: " + err.Error(),
		})
		return
	}

	// 故意使用不安全的XML解析方式
	var user User
	err = xml.Unmarshal(body, &user)
	if err != nil {
		c.HTML(http.StatusOK, "vulns/xxe.tmpl", gin.H{
			"title": "XXE漏洞演示",
			"error": "XML解析失败: " + err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "vulns/xxe.tmpl", gin.H{
		"title":    "XXE漏洞演示",
		"success":  "XML解析成功",
		"username": user.Name,
		"email":    user.Email,
	})
}