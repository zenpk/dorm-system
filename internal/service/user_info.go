package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/cookie"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/util"
	"net/http"
)

type UserInfo struct{}

var userInfoDal dal.UserInfo

// GetAllUserInfos 获取全部用户信息
func (u *UserInfo) GetAllUserInfos(c *gin.Context) {
	userInfos, err := userInfoDal.FindAll()
	if err != nil {
		// 记录并返回一些信息
		return
	}
	c.JSON(http.StatusOK, dto.UserInfoFindAllResp{
		UserInfos: userInfos,
	})
}

// GetUserInfo 获取当前登录用户的信息
func (u *UserInfo) GetUserInfo(c *gin.Context) {
	idStr, err := cookie.GetUserId(c)
	if err != nil {
		// 记录并返回一些信息
		return
	}
	id := util.ParseId(idStr)
	userInfo, err := userInfoDal.FindById(id)
	if err != nil {
		// 记录并返回一些信息
		return
	}
	c.JSON(http.StatusOK, dto.UserInfoResp{
		UserInfo: userInfo,
	})
}
