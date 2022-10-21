package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/dal"
	"github.com/zenpk/dorm-system/internal/dto"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/eh"
	"net/http"
)

type Building struct{}

func (b *Building) GetAvailableCount(c *gin.Context) {
	id := util.QueryU32(c, "buildingId")
	var table *dal.Dorm
	sum, err := table.SumAvailableByBuildingId(id)
	errHandler := eh.JSONHandler{C: c, V: dto.GetAvailableCountResp{}}
	if err != nil {
		errHandler.Handle(err)
		return
	}
	c.JSON(http.StatusOK, dto.GetAvailableCountResp{
		CommonResp: dto.CommonResp{
			Code: eh.Preset.CodeOK,
			Msg:  "success",
		},
		Count: sum,
	})
}

func (b *Building) GetAvailableBuildings(c *gin.Context) {
	var table *dal.Building
	buildings, err := table.FindAllAvailable()
	errHandler := eh.JSONHandler{C: c, V: dto.GetAvailableBuildingsResp{}}
	if err != nil {
		errHandler.Handle(err)
		return
	}
	var ids []int64
	for _, b := range buildings {
		ids = append(ids, int64(b.BuildingId))
	}
	c.JSON(http.StatusOK, dto.GetAvailableBuildingsResp{
		CommonResp: dto.CommonResp{
			Code: eh.Preset.CodeOK,
			Msg:  "success",
		},
		BuildingIds: ids,
	})
}
