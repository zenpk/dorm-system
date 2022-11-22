package controller

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/pkg/zap"
	"time"
)

func InitGin() error {
	gin.SetMode(viper.GetString("gin.mode"))
	r := gin.Default()
	if viper.GetString("gin.mode") == "release" {
		// Add a ginzap middleware, which:
		//   - Logs all requests, like a combined access and error log.
		//   - Logs to stdout.
		//   - RFC3339 with UTC time format.
		r.Use(ginzap.Ginzap(zap.Logger.Desugar(), time.RFC3339, true)) // enable zap on release mode
		// Logs all panic to error log
		//   - stack means whether output the stack info.
		r.Use(ginzap.RecoveryWithZap(zap.Logger.Desugar(), false))
	}
	InitRouter(r)
	addr := fmt.Sprintf("%s:%d", viper.GetString("server.host"), viper.GetInt("server.port"))
	zap.Logger.Infof("server listening at %v", addr)
	err := r.Run(addr)
	return err
}
