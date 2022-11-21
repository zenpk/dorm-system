package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dto"
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
}

func (t *Team) Get(c *gin.Context) {
	packer := ep.Packer{V: dto.TeamCreateGetResp{}}
	userIdStr := cookie.GetUserId(c)
	userId := util.ParseU64(userIdStr)
	if userId <= 0 { // userId shouldn't be 0
		response(c, packer.Pack(ep.ErrInputHeader))
	}
}
