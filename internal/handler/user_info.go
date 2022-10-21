package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/eh"
	"net/http"
)

type UserInfo struct{}

// GetMyInfo get UserInfo based on the id in Cookie
func (*UserInfo) GetMyInfo(c *gin.Context) {
	idStr, err := cookie.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, dto.GetUserInfoResp{
			CommonResp: dto.CommonResp{
				Code: eh.Preset.CodeTokenError,
				Msg:  "you're not logged in",
			},
		})
		return
	}
	id := util.ParseU64(idStr)
	var userInfo *dal.UserInfo
	userInfo, err = userInfo.FindById(id)
	errHandler := eh.JSONHandler{C: c, V: dto.GetUserInfoResp{}}
	if err != nil {
		errHandler.Handle(err)
		return
	}
	c.JSON(http.StatusOK, dto.GetUserInfoResp{
		CommonResp: dto.CommonResp{
			Code: eh.Preset.CodeOK,
			Msg:  "success",
		},
		UserInfo: userInfo,
	})
}
