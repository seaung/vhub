package api

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// SSRFHandler 处理SSRF漏洞演示页面
func SSRFHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/ssrf.tmpl", gin.H{
		"title": "SSRF漏洞演示",
	})
}

// SSRFRequestHandler 处理SSRF请求转发（不安全的方式）
func SSRFRequestHandler(c *gin.Context) {
	targetURL := c.Query("url")
	if targetURL == "" {
		c.HTML(http.StatusOK, "vulns/ssrf.tmpl", gin.H{
			"title": "SSRF漏洞演示",
			"error": "请提供目标URL",
		})
		return
	}

	// 故意不做URL验证，允许访问任意URL，包括内网地址
	client := &http.Client{}
	req, err := http.NewRequest("GET", targetURL, nil)
	if err != nil {
		c.HTML(http.StatusOK, "vulns/ssrf.tmpl", gin.H{
			"title": "SSRF漏洞演示",
			"error": "创建请求失败: " + err.Error(),
		})
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		c.HTML(http.StatusOK, "vulns/ssrf.tmpl", gin.H{
			"title": "SSRF漏洞演示",
			"error": "请求失败: " + err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.HTML(http.StatusOK, "vulns/ssrf.tmpl", gin.H{
			"title": "SSRF漏洞演示",
			"error": "读取响应失败: " + err.Error(),
		})
		return
	}

	// 获取响应头信息
	headers := make(map[string]string)
	for key, values := range resp.Header {
		headers[key] = strings.Join(values, ", ")
	}

	c.HTML(http.StatusOK, "vulns/ssrf.tmpl", gin.H{
		"title":    "SSRF漏洞演示",
		"success":  "请求成功",
		"url":      targetURL,
		"status":   resp.Status,
		"headers":  headers,
		"response": string(body),
	})
}