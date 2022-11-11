package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/rpc"
	pb "github.com/zenpk/dorm-system/internal/service/user"
	"github.com/zenpk/dorm-system/pkg/ep"
	"strconv"
)

type User struct{}

func (u *User) Register(c *gin.Context) {
	var req pb.RegisterLoginRequest
	packer := ep.Packer{V: dto.CommonResp{}}
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	// check password length
	if len(req.Password) < viper.GetInt("auth.password_length") {
		errPack := ep.ErrInputBody
		errPack.Msg = "password too short"
		response(c, packer.Pack(errPack))
		return
	}
	resp, err := rpc.Client.User.Register(&req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	cookie.SetToken(c, resp.Token)
	cookie.SetUserId(c, strconv.FormatUint(resp.UserId, 10))
	cookie.SetUsername(c, resp.Username)
	dtoResp := resp.Resp
	response(c, dtoResp)
}

func (u *User) Login(c *gin.Context) {
	var req pb.RegisterLoginRequest
	packer := ep.Packer{V: dto.CommonResp{}}
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	// check password length
	if len(req.Password) < viper.GetInt("auth.password_length") {
		errPack := ep.ErrInputBody
		errPack.Msg = "password too short"
		response(c, packer.Pack(errPack))
		return
	}
	resp, err := rpc.Client.User.Login(&req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	cookie.SetToken(c, resp.Token)
	cookie.SetUserId(c, strconv.FormatUint(resp.UserId, 10))
	cookie.SetUsername(c, resp.Username)
	dtoResp := resp.Resp
	response(c, dtoResp)
}

func (u *User) Logout(c *gin.Context) {
	cookie.ClearAllUserInfos(c)
	response(c, dto.CommonResp{
		Code: 0,
		Msg:  "successfully logged out",
	})
}

func (u *User) UpdatePassword(c *gin.Context) {

}
