package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/rpc"
	pb "github.com/zenpk/dorm-system/internal/service/order"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Order struct{}

func (o *Order) Submit(c *gin.Context) {
	var req pb.OrderRequest
	packer := ep.Packer{V: dto.CommonResp{}}
	if err := c.ShouldBindJSON(&req); err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	resp, err := rpc.Client.Order.Submit(&req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	dtoResp := resp.Resp
	response(c, dtoResp)
}
