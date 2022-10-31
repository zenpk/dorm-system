package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/rpc"
	"github.com/zenpk/dorm-system/pkg/ep"
)

type Dorm struct{}

func (d *Dorm) GetAvailableNum(c *gin.Context) {
	packer := &ep.Packer{V: dto.CommonResp{}}
	resp, err := rpc.Client.Dorm.GetAvailableNum(nil)
	if err != nil {
		response(c, packer.PackWithError(err))
		return
	}
	response(c, resp)
}
