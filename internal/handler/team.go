package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/rpc"
	pb "github.com/zenpk/dorm-system/internal/service/team"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Team struct{}

func (t *Team) Create(c *gin.Context) {
	packer := ep.Packer{V: dto.CommonResp{}}
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	if userId <= 0 { // userId shouldn't be 0
		response(c, packer.Pack(ep.ErrInputHeader))
	}
	req := &pb.CreateGetRequest{UserId: userId}
	resp, err := rpc.Client.Team.Create(req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}

func (t *Team) Get(c *gin.Context) {
	packer := ep.Packer{V: dto.CommonResp{}}
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	if userId <= 0 { // userId shouldn't be 0
		response(c, packer.Pack(ep.ErrInputHeader))
	}
	req := &pb.CreateGetRequest{UserId: userId}
	resp, err := rpc.Client.Team.Get(req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}

func (t *Team) Join(c *gin.Context) {
	packer := ep.Packer{V: dto.CommonResp{}}
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	if userId <= 0 { // userId shouldn't be 0
		response(c, packer.Pack(ep.ErrInputHeader))
	}
	var req pb.JoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, packer.Pack(ep.ErrInputBody))
	}
	req.UserId = userId
	resp, err := rpc.Client.Team.Join(&req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}
