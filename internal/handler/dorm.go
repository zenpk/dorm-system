package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/rpc"
	pb "github.com/zenpk/dorm-system/internal/service/dorm"
)

type Dorm struct{}

func (d Dorm) GetRemainCnt(c *gin.Context) {
	req := new(pb.EmptyRequest)
	resp, err := rpc.Client.Dorm.GetRemainCnt(req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}

func (d Dorm) GetAll(c *gin.Context) {
	req := new(pb.EmptyRequest)
	resp, err := rpc.Client.Dorm.GetAll(req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}
