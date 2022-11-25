package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/handler"
	"github.com/zenpk/dorm-system/internal/util"
	"github.com/zenpk/dorm-system/pkg/ep"
	"net/http"
	"time"
)

func CheckUserTeamTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		packer := ep.Packer{V: handler.CommonResp{}}
		startTime, err := util.StringToTime(viper.GetString("datetime.user_team_start_time"))
		if err != nil {
			c.JSON(http.StatusOK, packer.PackWithError(err))
			c.Abort()
			return
		}
		deadline, err := util.StringToTime(viper.GetString("datetime.user_team_deadline"))
		if err != nil {
			c.JSON(http.StatusOK, packer.PackWithError(err))
			c.Abort()
			return
		}
		if time.Now().Before(startTime) {
			errPack := ep.ErrLogic
			errPack.Msg = "not yet user/team modification time"
			c.JSON(http.StatusOK, packer.Pack(err))
			c.Abort()
			return
		}
		if time.Now().After(deadline) {
			errPack := ep.ErrLogic
			errPack.Msg = "user/team modification deadline is exceeded"
			c.JSON(http.StatusOK, packer.Pack(err))
			c.Abort()
			return
		}
		c.Next()
	}
}

func CheckOrderTime() gin.HandlerFunc {
	return func(c *gin.Context) {
		packer := ep.Packer{V: handler.CommonResp{}}
		startTime, err := util.StringToTime(viper.GetString("datetime.order_start_time"))
		if err != nil {
			c.JSON(http.StatusOK, packer.PackWithError(err))
			c.Abort()
			return
		}
		deadline, err := util.StringToTime(viper.GetString("datetime.order_deadline"))
		if err != nil {
			c.JSON(http.StatusOK, packer.PackWithError(err))
			c.Abort()
			return
		}
		if time.Now().Before(startTime) {
			errPack := ep.ErrLogic
			errPack.Msg = "not yet order modification start time"
			c.JSON(http.StatusOK, packer.Pack(err))
			c.Abort()
			return
		}
		if time.Now().After(deadline) {
			errPack := ep.ErrLogic
			errPack.Msg = "order modification deadline is exceeded"
			c.JSON(http.StatusOK, packer.Pack(err))
			c.Abort()
			return
		}
		c.Next()
	}
}
