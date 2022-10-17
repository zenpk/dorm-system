package eh

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/pkg/zap"
	"net/http"
)

type Building struct {
	C *gin.Context
}

func (b *Building) GetAvailableCountErr(err error) {
	zap.Logger.Warn(err.Error())
	b.C.JSON(http.StatusOK, dto.GetAvailableCountResp{
		CommonResp: dto.CommonResp{
			Code: CodeUncaughtError,
			Msg:  err.Error(),
		},
		Count: 0,
	})
}

func (b *Building) GetAvailableBuildingsErr(err error) {
	zap.Logger.Warn(err.Error())
	b.C.JSON(http.StatusOK, dto.GetAvailableBuildingsResp{
		CommonResp: dto.CommonResp{
			Code: CodeUncaughtError,
			Msg:  err.Error(),
		},
		BuildingIds: nil,
	})
}
