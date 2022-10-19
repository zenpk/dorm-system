package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/pkg/eh"
	"github.com/zenpk/dorm-system/pkg/zap"
	"net/http"
)

// CheckAuthInfo extract user infos from the JWT token in Cookie
func CheckAuthInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := cookie.GetToken(c)
		if err != nil || token == "" { // no Cookie
			c.Next()
			return
		}
		userId, username, err := ParseToken(token)
		if err != nil || userId == "0" {
			zap.Logger.Warn("ParseToken failed")
			c.JSON(http.StatusUnauthorized, dto.CommonResp{
				Code: eh.Preset.CodeMiddlewareError,
				Msg:  "ParseToken failed",
			})
			return
		}
		cookie.SetUserId(c, userId)
		cookie.SetUsername(c, username)
		c.Next()
	}
}

// RequireLogin if not login then abort
func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := cookie.GetToken(c)
		if err != nil || token == "" { // no Cookie
			c.JSON(http.StatusUnauthorized, dto.CommonResp{
				Code: eh.Preset.CodeMiddlewareError,
				Msg:  "login required",
			})
			c.Abort()
			return
		}
		userId, username, err := ParseToken(token)
		if err != nil || userId == "0" {
			zap.Logger.Warn("ParseToken failed")
			c.JSON(http.StatusUnauthorized, dto.CommonResp{
				Code: eh.Preset.CodeMiddlewareError,
				Msg:  "ParseToken failed",
			})
			return
		}
		cookie.SetUserId(c, userId)
		cookie.SetUsername(c, username)
		c.Next()
	}
}
