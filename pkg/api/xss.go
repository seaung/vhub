package api

import (
	"github.com/gin-gonic/gin"
	"github.com/seaung/vhub/pkg/models"
	"net/http"
	"time"
)

// XSSDOMHandler 处理DOM型XSS演示页面
func XSSDOMHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/xss_dom.tmpl", gin.H{})
}

// XSSReflectedHandler 处理反射型XSS演示页面
func XSSReflectedHandler(c *gin.Context) {
	query := c.Query("q")
	c.HTML(http.StatusOK, "vulns/xss_reflected.tmpl", gin.H{
		"Query": query,
	})
}

// XSSStoredHandler 处理存储型XSS演示页面的GET请求
func XSSStoredHandler(c *gin.Context) {
	messages, err := models.GetAllMessages()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error loading messages")
		return
	}

	c.HTML(http.StatusOK, "vulns/xss_stored.tmpl", gin.H{
		"messages": messages,
	})
}

// XSSStoredSubmitHandler 处理存储型XSS演示页面的POST请求
func XSSStoredSubmitHandler(c *gin.Context) {
	username := c.PostForm("username")
	content := c.PostForm("message")

	message := &models.Message{
		Username:  username,
		Content:   content,
		CreatedAt: time.Now(),
	}

	if err := models.CreateMessage(message); err != nil {
		c.String(http.StatusInternalServerError, "Error saving message")
		return
	}

	c.Redirect(http.StatusFound, "/xss/stored")
}