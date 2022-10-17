package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/eh"
	"github.com/zenpk/dorm-system/internal/util"
	"net/http"
)

type UserInfo struct{}

// GetMyInfo get UserInfo based on the id in Cookie
func (*UserInfo) GetMyInfo(c *gin.Context) {
	idStr, err := cookie.GetUserId(c)
	if err != nil {
		c.JSON(http.StatusOK, dto.GetUserInfoResp{
			CommonResp: dto.CommonResp{
				Code: eh.CodeTokenError,
				Msg:  "you're not logged in",
			},
			UserInfo: dal.UserInfo{},
		})
		return
	}
	id := util.ParseU64(idStr)
	var userInfo dal.UserInfo
	userInfo, err = userInfo.FindById(id)
	errHandler := eh.User{C: c}
	if err != nil {
		errHandler.GetMyInfoErr(err)
		return
	}
	c.JSON(http.StatusOK, dto.GetUserInfoResp{
		CommonResp: dto.CommonResp{
			Code: eh.CodeOK,
			Msg:  "success",
		},
		UserInfo: userInfo,
	})
}
