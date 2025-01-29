package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Username string
	Password string
	Balance  float64
}

var users = map[string]User{
	"admin": {Username: "admin", Password: "admin123", Balance: 1000.0},
}

func CSRFHandler(c *gin.Context) {
	username, _ := c.Cookie("username")
	c.HTML(http.StatusOK, "vulns/csrf.tmpl", gin.H{
		"isLoggedIn": username != "",
		"username":   username,
	})
}

func CSRFLoginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	if user, exists := users[username]; exists && user.Password == password {
		c.SetCookie("username", username, 3600, "/", "", false, true)
		c.Redirect(http.StatusFound, "/csrf")
		return
	}

	c.HTML(http.StatusOK, "vulns/csrf.tmpl", gin.H{
		"error":   true,
		"message": "用户名或密码错误",
	})
}

func CSRFLogoutHandler(c *gin.Context) {
	c.SetCookie("username", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, "/csrf")
}

func CSRFChangePasswordHandler(c *gin.Context) {
	username, _ := c.Cookie("username")
	if username == "" {
		c.Redirect(http.StatusFound, "/csrf")
		return
	}

	newPassword := c.PostForm("newPassword")
	if user, exists := users[username]; exists {
		user.Password = newPassword
		users[username] = user
		c.HTML(http.StatusOK, "vulns/csrf.tmpl", gin.H{
			"isLoggedIn": true,
			"username":   username,
			"message":    "密码修改成功",
		})
		return
	}

	c.HTML(http.StatusOK, "vulns/csrf.tmpl", gin.H{
		"error":   true,
		"message": "修改密码失败",
	})
}

func CSRFTransferHandler(c *gin.Context) {
	username, _ := c.Cookie("username")
	if username == "" {
		c.Redirect(http.StatusFound, "/csrf")
		return
	}

	to := c.PostForm("to")
	amount := c.PostForm("amount")

	c.HTML(http.StatusOK, "vulns/csrf.tmpl", gin.H{
		"isLoggedIn": true,
		"username":   username,
		"message":    "转账成功：已向 " + to + " 转账 " + amount + " 元",
	})
}