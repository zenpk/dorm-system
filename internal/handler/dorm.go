package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/rpc"
	pb "github.com/zenpk/dorm-system/internal/service/dorm"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Dorm struct{}

func (d *Dorm) GetAvailableNum(c *gin.Context) {
	packer := &ep.Packer{V: dto.CommonResp{}}
	req := new(pb.EmptyRequest)
	resp, err := rpc.Client.Dorm.GetAvailableNum(req)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}
