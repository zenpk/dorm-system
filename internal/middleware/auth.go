package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/pkg/ep"
	"net/http"
)

// CheckAuthInfo extract user infos from the JWT token in Cookie
func CheckAuthInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := cookie.GetToken(c)
		if err != nil || token == "" { // no cookie
			c.Next()
			return
		}
		if err := cookie.SetAllFromToken(c, token); err != nil {
			packer := ep.Packer{V: dto.CommonResp{}}
			errPack := ep.ErrInputToken
			c.JSON(http.StatusUnauthorized, packer.PackWithError(errPack))
			c.Abort()
			return
		}
		c.Next()
	}
}

// RequireLogin if no token then abort
func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := cookie.GetToken(c)
		packer := ep.Packer{V: dto.CommonResp{}}
		if err != nil || token == "" { // no cookie
			errPack := ep.ErrNotLogin
			c.JSON(http.StatusUnauthorized, packer.Pack(errPack))
			c.Abort()
			return
		}
		if err := cookie.SetAllFromToken(c, token); err != nil {
			errPack := ep.ErrNotLogin
			errPack.Msg = ep.ErrInputToken.Msg
			c.JSON(http.StatusUnauthorized, packer.PackWithError(errPack))
			c.Abort()
			return
		}
		c.Next()
	}
}
