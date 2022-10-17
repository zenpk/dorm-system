package eh

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/middleware"
	"net/http"
)

type User struct {
	C *gin.Context
}

func (u *User) RegisterLoginErr(err error) {
	middleware.Logger.Warn(err.Error())
	u.C.JSON(http.StatusOK, dto.RegisterLoginResp{
		SuccessCode: 0,
		Status:      -1,
		Message:     err.Error(),
		Data:        dal.UserInfo{},
	})
}
