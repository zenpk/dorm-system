package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/pkg/ep"
	"net/http"
)

type CommonResp struct {
	Err ep.ErrPack `json:"err,omitempty"`
}

type finalResponse struct {
	v interface{}
}

var packer = ep.Packer{V: CommonResp{}}

func response(c *gin.Context, v interface{}) {
	c.JSON(http.StatusOK, v)
}
