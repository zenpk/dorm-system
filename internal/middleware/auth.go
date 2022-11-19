package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/rpc"
	pb "github.com/zenpk/dorm-system/internal/service/token"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/ep"
	"net/http"
)

// CheckAuthInfo extract user infos from the JWT token in Cookie, won't abort if not logged in
func CheckAuthInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		packer := ep.Packer{V: dto.CommonResp{}}
		token := cookie.GetAccessToken(c)
		if token == "" { // no access_token
			refreshToken := cookie.GetRefreshToken(c)
			if refreshToken != "" {
				req := &pb.GenAccessTokenRequest{RefreshToken: refreshToken}
				tokenResp, err := rpc.Client.Token.GenAccessToken(req)
				if err != nil {
					c.JSON(http.StatusOK, packer.PackWithError(err))
					c.Abort()
					return
				}
				accessToken := tokenResp.AccessToken
				cookie.SetAccessToken(c, accessToken)
				if err := cookie.SetAllFromAccessToken(c, accessToken); err != nil {
					c.JSON(http.StatusOK, packer.PackWithError(err))
					c.Abort()
					return
				}
			}
		} else if err := cookie.SetAllFromAccessToken(c, token); err != nil {
			c.JSON(http.StatusOK, packer.PackWithError(err))
			c.Abort()
			return
		}
		c.Next()
	}
}

// RequireLogin if no token then abort
func RequireLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := cookie.GetAccessToken(c)
		packer := ep.Packer{V: dto.CommonResp{}}
		if token == "" { // no access_token
			refreshToken := cookie.GetRefreshToken(c)
			if refreshToken == "" { // no refresh_token
				c.JSON(http.StatusOK, packer.Pack(ep.ErrNotLogin))
				c.Abort()
				return
			}
			req := &pb.GenAccessTokenRequest{RefreshToken: refreshToken}
			tokenResp, err := rpc.Client.Token.GenAccessToken(req)
			if err != nil {
				c.JSON(http.StatusOK, packer.PackWithError(err))
				c.Abort()
				return
			}
			accessToken := tokenResp.AccessToken
			cookie.SetAccessToken(c, accessToken)
			if err := cookie.SetAllFromAccessToken(c, accessToken); err != nil {
				c.JSON(http.StatusOK, packer.PackWithError(err))
				c.Abort()
				return
			}
		} else if err := cookie.SetAllFromAccessToken(c, token); err != nil {
			c.JSON(http.StatusOK, packer.PackWithError(err))
			c.Abort()
			return
		}
		c.Next()
	}
}

// RequireAdmin must be logged in and has admin role
func RequireAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := cookie.GetAccessToken(c)
		packer := ep.Packer{V: dto.CommonResp{}}
		if token == "" { // no access_token
			refreshToken := cookie.GetRefreshToken(c)
			if refreshToken == "" { // no refresh_token
				c.JSON(http.StatusOK, packer.Pack(ep.ErrNotLogin))
				c.Abort()
				return
			}
			req := &pb.GenAccessTokenRequest{RefreshToken: refreshToken}
			tokenResp, err := rpc.Client.Token.GenAccessToken(req)
			if err != nil {
				c.JSON(http.StatusOK, packer.PackWithError(err))
				c.Abort()
				return
			}
			accessToken := tokenResp.AccessToken
			cookie.SetAccessToken(c, accessToken)
			if err := cookie.SetAllFromAccessToken(c, accessToken); err != nil {
				c.JSON(http.StatusOK, packer.PackWithError(err))
				c.Abort()
				return
			}
		} else if err := cookie.SetAllFromAccessToken(c, token); err != nil {
			c.JSON(http.StatusOK, packer.PackWithError(err))
			c.Abort()
			return
		}
		// check admin role
		roleStr := cookie.GetRole(c)
		role := util.Parse32(roleStr)
		if role <= 1 { // not admin
			c.JSON(http.StatusOK, packer.Pack(ep.ErrNoPermission))
			c.Abort()
			return
		}
		c.Next()
	}
}
