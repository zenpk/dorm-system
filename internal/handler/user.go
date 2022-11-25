package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/rpc"
	tokenpb "github.com/zenpk/dorm-system/internal/service/token"
	userpb "github.com/zenpk/dorm-system/internal/service/user"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type User struct{}

func (u User) Register(c *gin.Context) {
	var userReq userpb.RegisterLoginRequest
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
	response(c, packer.Pack(ep.ErrOK))
}

func (u User) Login(c *gin.Context) {
	var userReq userpb.RegisterLoginRequest
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
	response(c, packer.Pack(ep.ErrOK))
}

func (u User) Logout(c *gin.Context) {
	cookie.ClearAllUserInfos(c)
	errPack := ep.ErrOK
	errPack.Msg = "successfully logged out"
	response(c, CommonResp{
		Err: errPack,
	})
}

func (u User) Get(c *gin.Context) {
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	if userId <= 0 { // userId shouldn't be 0
		response(c, packer.Pack(ep.ErrInputHeader))
		return
	}
	req := &userpb.GetRequest{UserId: userId}
	resp, err := rpc.Client.User.Get(req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}

func (u User) Edit(c *gin.Context) {
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	if userId <= 0 { // userId shouldn't be 0
		response(c, packer.Pack(ep.ErrInputHeader))
		return
	}
	var req userpb.EditRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	// write in userId
	req.User.Id = userId
	resp, err := rpc.Client.User.Edit(&req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}
