package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/mq"
	"github.com/zenpk/dorm-system/internal/rpc"
	pb "github.com/zenpk/dorm-system/internal/service/order"
	"github.com/zenpk/dorm-system/internal/service/team"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Order struct{}

func (o Order) Submit(c *gin.Context) {
	var req pb.SubmitRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	if len(req.BuildingNum) <= 0 {
		errPack := ep.ErrInputBody
		errPack.Msg = "wrong building number"
		response(c, packer.Pack(errPack))
		return
	}
	// First: get teamId
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	teamReq := &team.GetRequest{UserId: userId}
	teamResp, err := rpc.Client.Team.Get(teamReq)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	// Second: check if user is the owner
	if teamResp.Team.Owner.Id != userId {
		errPack := ep.ErrNoPermission
		errPack.Msg = "you're not the team owner"
		response(c, packer.Pack(errPack))
		return
	}
	req.Team = teamResp.Team
	// Third: generate a unique code
	code, err := util.GenSnowflakeString()
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	req.Code = code
	if err := mq.Producer.Order.Send(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	errPack := ep.ErrOK
	errPack.Msg = "submitted successfully"
	response(c, packer.Pack(errPack))
}

func (o Order) Get(c *gin.Context) {
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	teamReq := &team.GetRequest{UserId: userId}
	teamResp, err := rpc.Client.Team.Get(teamReq)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	req := &pb.GetRequest{TeamId: teamResp.Team.Id}
	resp, err := rpc.Client.Order.Get(req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}

func (o Order) Delete(c *gin.Context) {
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	orderId := util.QueryU64(c, "orderId")
	if orderId <= 0 { // orderId shouldn't be 0
		response(c, packer.Pack(ep.ErrInputBody))
		return
	}
	teamReq := &team.GetRequest{UserId: userId}
	teamResp, err := rpc.Client.Team.Get(teamReq)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	// check if user is the owner
	if teamResp.Team.Owner.Id != userId {
		errPack := ep.ErrNoPermission
		errPack.Msg = "you're not the team owner"
		response(c, packer.Pack(errPack))
		return
	}
	var req pb.DeleteRequest
	req.OrderId = orderId
	req.Team = teamResp.Team
	resp, err := rpc.Client.Order.Delete(&req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}
