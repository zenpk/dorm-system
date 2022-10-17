package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
	"time"
)

// CORSFilter CORS 管理
func CORSFilter() gin.HandlerFunc {
	return cors.New(cors.Config{
		//AllowOrigins:     []string{"https://localhost:5173", "http://localhost:5173"}, // ignored because of AllowOriginFunc
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// 服务器的请求
			if strings.Contains(origin, viper.GetString("server.domain")) {
				return true
			}
			// 白名单的请求
			whitelist := viper.GetStringSlice("cors.whitelist")
			for _, allow := range whitelist {
				if strings.Contains(origin, allow) {
					return true
				}
			}
			return false
		},
		MaxAge: 12 * time.Hour,
	})
}
