package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/rpc"
	tokenpb "github.com/zenpk/dorm-system/internal/service/token"
	userpb "github.com/zenpk/dorm-system/internal/service/user"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type User struct{}

func (u *User) Register(c *gin.Context) {
	var userReq userpb.RegisterLoginRequest
	packer := ep.Packer{V: dto.CommonResp{}}
	if err := c.ShouldBindJSON(&userReq); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	// check password length
	if len(userReq.Password) < viper.GetInt("auth.password_length") {
		errPack := ep.ErrInputBody
		errPack.Msg = "password too short"
		response(c, packer.Pack(errPack))
		return
	}
	userResp, err := rpc.Client.User.Register(&userReq)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	tokenReq := &tokenpb.GenAllTokenRequest{UserId: userResp.UserId}
	tokenResp, err := rpc.Client.Token.GenAllToken(tokenReq)
	cookie.SetAccessToken(c, tokenResp.AccessToken)
	cookie.SetRefreshToken(c, tokenResp.RefreshToken)
	if err := cookie.SetAllFromAccessToken(c, tokenResp.AccessToken); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, tokenResp.Resp)
}

func (u *User) Login(c *gin.Context) {
	var userReq userpb.RegisterLoginRequest
	packer := ep.Packer{V: dto.CommonResp{}}
	if err := c.ShouldBindJSON(&userReq); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	// check password length
	if len(userReq.Password) < viper.GetInt("auth.password_length") {
		errPack := ep.ErrInputBody
		errPack.Msg = "password too short"
		response(c, packer.Pack(errPack))
		return
	}
	userResp, err := rpc.Client.User.Login(&userReq)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	tokenReq := &tokenpb.GenAllTokenRequest{UserId: userResp.UserId}
	tokenResp, err := rpc.Client.Token.GenAllToken(tokenReq)
	cookie.SetAccessToken(c, tokenResp.AccessToken)
	cookie.SetRefreshToken(c, tokenResp.RefreshToken)
	if err := cookie.SetAllFromAccessToken(c, tokenResp.AccessToken); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, tokenResp.Resp)
}

func (u *User) Logout(c *gin.Context) {
	cookie.ClearAllUserInfos(c)
	response(c, dto.CommonResp{
		Code: ep.ErrOK.Code,
		Msg:  "successfully logged out",
	})
}

func (u *User) UpdatePassword(c *gin.Context) {

}
