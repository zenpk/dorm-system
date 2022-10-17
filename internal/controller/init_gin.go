package controller

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/zenpk/dorm-system/internal/middleware"
	"time"
)

// InitGin 初始化 Gin
func InitGin() error {
	gin.SetMode(viper.GetString("gin.mode")) // 设置运行模式
	r := gin.Default()
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	if viper.GetString("gin.mode") == "release" {
		r.Use(ginzap.Ginzap(middleware.Logger, time.RFC3339, true)) // release 模式下启用 Logger 中间件
	}
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(middleware.Logger, true))
	InitRouter(r)
	err := r.Run(viper.GetString("server.host") + ":" + viper.GetString("server.port")) // 开始监听
	return err
}
