package eh

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/pkg/zap"
	"net/http"
)

type User struct {
	C *gin.Context
}

func (u *User) RegisterLoginErr(err error) {
	zap.Logger.Warn(err.Error())
	u.C.JSON(http.StatusOK, dto.RegisterLoginResp{
		Code:     CodeUncaughtError,
		Msg:      err.Error(),
		UserInfo: dal.UserInfo{},
	})
}
