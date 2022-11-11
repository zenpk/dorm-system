package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/mq"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Order struct{}

func (o *Order) Submit(c *gin.Context) {
	var req dto.OrderRequest
	packer := ep.Packer{V: dto.CommonResp{}}
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	if len(req.StudentNum1) <= 0 {
		errPack := ep.ErrInputBody
		errPack.Msg = "studentNum1 mustn't be empty"
		response(c, packer.PackWithInfo(errPack))
		return
	}
	if err := mq.Producer.Order.Send(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	errPack := ep.ErrOK
	errPack.Msg = "submitted successfully"
	response(c, packer.Pack(errPack))
}
