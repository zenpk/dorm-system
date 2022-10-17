package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"net/http"
)

func HelloDemo(c *gin.Context) {
	c.String(http.StatusOK, "Hello!")
}

func UserDemo(c *gin.Context) {
	c.String(http.StatusOK, "You've logged In!")
}

func AdminDemo(c *gin.Context) {
	role, _ := cookie.GetRole(c)
	if role != "ADMIN" {
		c.String(http.StatusUnauthorized, "no permission!")
		return
	}
	c.String(http.StatusOK, "You are the admin!")
}
