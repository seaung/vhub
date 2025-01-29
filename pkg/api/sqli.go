package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/seaung/vhub/pkg/models"
)

// SQLIHandler 展示SQL注入漏洞演示页面
func SQLIHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "vulns/sqli.tmpl", gin.H{
		"title": "SQL注入漏洞演示",
	})
}

// SQLILoginHandler 演示登录验证中的SQL注入
func SQLILoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// 故意使用字符串拼接构造SQL语句，易受SQL注入攻击
	query := fmt.Sprintf("SELECT id, name FROM users WHERE name = '%s' AND password = '%s'", username, password)

	var user models.Users
	row := models.DB.Raw(query).Row()
	err := row.Scan(&user.ID, &user.Name)

	if err == sql.ErrNoRows {
		c.HTML(http.StatusOK, "vulns/sqli.tmpl", gin.H{
			"error": "用户名或密码错误",
		})
		return
	}

	if err != nil {
		c.HTML(http.StatusOK, "vulns/sqli.tmpl", gin.H{
			"error": "查询出错: " + err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "vulns/sqli.tmpl", gin.H{
		"success": true,
		"message": fmt.Sprintf("欢迎回来，%s!", user.Name),
	})
}

// SQLISearchHandler 演示消息搜索中的SQL注入
func SQLISearchHandler(c *gin.Context) {
	keyword := c.Query("keyword")

	// 故意使用LIKE语句和字符串拼接，易受SQL注入攻击
	query := fmt.Sprintf("SELECT id, content FROM messages WHERE content LIKE '%%%s%%'", keyword)

	rows, err := models.DB.Raw(query).Rows()
	if err != nil {
		c.HTML(http.StatusOK, "vulns/sqli.tmpl", gin.H{
			"error": "搜索出错: " + err.Error(),
		})
		return
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.ID, &msg.Content); err != nil {
			continue
		}
		messages = append(messages, msg)
	}

	c.HTML(http.StatusOK, "vulns/sqli.tmpl", gin.H{
		"messages": messages,
		"keyword":  keyword,
	})
}

// SQLIUpdateHandler 演示用户信息更新中的SQL注入
func SQLIUpdateHandler(c *gin.Context) {
	userID := c.PostForm("user_id")
	description := c.PostForm("description")

	// 故意使用字符串拼接构造UPDATE语句，易受SQL注入攻击
	query := fmt.Sprintf("UPDATE users SET description = '%s' WHERE id = %s", description, userID)

	result := models.DB.Exec(query)
	if err := result.Error; err != nil {
		c.HTML(http.StatusOK, "vulns/sqli.tmpl", gin.H{
			"error": "更新失败: " + err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "vulns/sqli.tmpl", gin.H{
		"success": true,
		"message": "用户信息已更新",
	})
}