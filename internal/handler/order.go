package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/mq"
	pb "github.com/zenpk/dorm-system/internal/service/order"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Order struct{}

func (o *Order) Submit(c *gin.Context) {
	var req dto.OrderReqSubmit
	packer := ep.Packer{V: dto.CommonResp{}}
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	// TODO
	reqPb := &pb.SubmitRequest{
		BuildingNum: req.BuildingNum,
		TeamId:      0,
	}
	if err := mq.Producer.Order.Send(reqPb); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	errPack := ep.ErrOK
	errPack.Msg = "submitted successfully"
	response(c, packer.Pack(errPack))
}

func (o *Order) Get(c *gin.Context) {
	return
}
