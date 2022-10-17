package service

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/eh"
	"github.com/zenpk/dorm-system/internal/util"
	"net/http"
)

type Building struct{}

func (b *Building) GetAvailableCount(c *gin.Context) {
	id := util.QueryU32(c, "buildingId")
	var dorm dal.Dorm
	sum, err := dorm.SumAvailableByBuildingId(id)
	errHandler := eh.Building{C: c}
	if err != nil {
		errHandler.GetAvailableCountErr(err)
		return
	}
	c.JSON(http.StatusOK, dto.GetAvailableCountResp{
		CommonResp: dto.CommonResp{
			Code: eh.CodeOK,
			Msg:  "success",
		},
		Count: sum,
	})
}

func (b *Building) GetAvailableBuildings(c *gin.Context) {
	var dalBuilding dal.Building
	buildings, err := dalBuilding.FindAllAvailable()
	errHandler := eh.Building{C: c}
	if err != nil {
		errHandler.GetAvailableBuildingsErr(err)
		return
	}
	var ids []int64
	for _, b := range buildings {
		ids = append(ids, int64(b.BuildingId))
	}
	c.JSON(http.StatusOK, dto.GetAvailableBuildingsResp{
		CommonResp: dto.CommonResp{
			Code: eh.CodeOK,
			Msg:  "success",
		},
		BuildingIds: ids,
	})
}
