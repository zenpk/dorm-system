package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/rpc"
	pb "github.com/zenpk/dorm-system/internal/service/team"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Team struct{}

func (t Team) Create(c *gin.Context) {
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	req := &pb.CreateRequest{UserId: userId}
	resp, err := rpc.Client.Team.Create(req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}

func (t Team) Get(c *gin.Context) {
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	req := &pb.GetRequest{UserId: userId}
	resp, err := rpc.Client.Team.Get(req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}

func (t Team) Join(c *gin.Context) {
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	var req pb.JoinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, packer.Pack(ep.ErrInputBody))
		return
	}
	req.UserId = userId
	resp, err := rpc.Client.Team.Join(&req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}

func (t Team) Leave(c *gin.Context) {
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	var req pb.LeaveRequest
	req.UserId = userId
	resp, err := rpc.Client.Team.Leave(&req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}

func (t Team) Transfer(c *gin.Context) {
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	var req pb.TransferRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	req.OldOwnerId = userId
	resp, err := rpc.Client.Team.Transfer(&req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}
