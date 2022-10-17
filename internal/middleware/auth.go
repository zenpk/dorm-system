package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"net/http"
)

// CheckAuthInfo 根据用户 Cookie 中的 token 信息提取 userId 和 username
func CheckAuthInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := cookie.GetToken(c)
		if err != nil || token == "" { // no cookie
			c.Next()
			return
		}
		userId, username, err := ParseToken(token)
		if err != nil || userId == "0" {
			Logger.Warn("ParseToken failed")
			return
		}
		cookie.SetUserId(c, userId)
		cookie.SetUsername(c, username)
		c.Next()
	}
}

// RequireLogin 需要登录的中间件，与上面的区别是如果未登录则直接 Abort
func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := cookie.GetToken(c)
		if err != nil || token == "" { // no cookie
			c.String(http.StatusUnauthorized, "you need to login")
			c.Abort()
			return
		}
		userId, username, err := ParseToken(token)
		if err != nil || userId == "0" {
			Logger.Warn("ParseToken failed")
			return
		}
		cookie.SetUserId(c, userId)
		cookie.SetUsername(c, username)
		c.Next()
	}
}
