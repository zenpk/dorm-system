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
	// First: get teamId
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	if userId <= 0 { // userId shouldn't be 0
		response(c, packer.Pack(ep.ErrInputHeader))
		return
	}
	teamReq := &team.CreateGetRequest{UserId: userId}
	teamResp, err := rpc.Client.Team.Get(teamReq)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	// Second: check if user is the owner
	if teamResp.Owner.Id != userId {
		errPack := ep.ErrNoPermission
		errPack.Msg = "you're not the team owner"
		response(c, packer.Pack(errPack))
		return
	}
	// Third: submit order
	req.Team = teamResp
	if err := mq.Producer.Order.Send(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	errPack := ep.ErrOK
	errPack.Msg = "submitted successfully"
	response(c, packer.Pack(errPack))
}

func (o Order) Get(c *gin.Context) {
	return
}
