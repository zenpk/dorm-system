package controller

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/zap"
	"time"
)

func InitGin() error {
	gin.SetMode(viper.GetString("gin.mode"))
	r := gin.Default()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	if viper.GetString("gin.mode") == "release" {
		r.Use(ginzap.Ginzap(zap.Logger, time.RFC3339, true)) // enable zap on release mode
	}
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(zap.Logger, true))
	InitRouter(r)
	err := r.Run(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
	return err
}
