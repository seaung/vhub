package api

import (
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	SafeDir = "./static/files/" // 安全的文件目录
)

// FileInclusion 处理文件包含漏洞演示页面
func FileInclusion(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/file_inclusion.html", gin.H{
		"title": "文件包含漏洞演示",
	})
}

// HandleLFI 处理本地文件包含
func HandleLFI(c *gin.Context) {
	fileName := c.Query("file")
	// 故意不做路径验证，造成任意文件读取漏洞
	filePath := filepath.Join(SafeDir, fileName)

	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		c.HTML(http.StatusOK, "vulns/file_inclusion.html", gin.H{
			"LFIContent": "Error: " + err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "vulns/file_inclusion.html", gin.H{
		"LFIContent": string(content),
	})
}

// HandleRFI 处理远程文件包含
func HandleRFI(c *gin.Context) {
	url := c.Query("url")
	// 故意不验证URL，允许包含任意远程文件
	resp, err := http.Get(url)
	if err != nil {
		c.HTML(http.StatusOK, "vulns/file_inclusion.html", gin.H{
			"RFIContent": "Error: " + err.Error(),
		})
		return
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.HTML(http.StatusOK, "vulns/file_inclusion.html", gin.H{
			"RFIContent": "Error: " + err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "vulns/file_inclusion.html", gin.H{
		"RFIContent": string(content),
	})
}