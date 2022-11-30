package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/middleware"
	"github.com/zenpk/dorm-system/test"
)

func InitRouter(router *gin.Engine) {
	router.Static("/static", "./assets/public") // static resources
	router.Use(middleware.CORSFilter())         // CORS
	// for temporary testing
	router.POST("/test-setup", test.SetupHandler)
	router.POST("/test", test.Handler)

	// no middleware URL
	router.POST("/login", ginHandler.user.Login)
	router.POST("/register", ginHandler.user.Register)
	router.GET("/remain-cnt", ginHandler.dorm.GetRemainCnt) // get the number of remaining beds count in each building
	router.GET("/buildings", ginHandler.dorm.GetAll)

	// login required URL
	auth := router.Group("/")
	auth.Use(middleware.RequireLogin())
	{
		user := auth.Group("/user")
		{
			user.GET("/info", ginHandler.user.Get)
			user.GET("/logout", ginHandler.user.Logout)
			user.Use(middleware.CheckUserTeamTime())
			{
				user.PUT("/info", ginHandler.user.Edit)
			}
		}
		team := auth.Group("/team")
		{
			team.GET("", ginHandler.team.Get) // get one's team info
			team.Use(middleware.CheckUserTeamTime())
			{
				team.POST("/create", ginHandler.team.Create)
				team.POST("/join", ginHandler.team.Join)
				team.DELETE("/leave", ginHandler.team.Leave)
			}
		}
		order := auth.Group("/order")
		{
			order.GET("", ginHandler.order.Get) // get one's team's order info
			order.Use(middleware.CheckOrderTime())
			{
				order.POST("/create", ginHandler.order.Submit)
				order.DELETE("/delete", ginHandler.order.Delete)
			}
		}
	}
	// admin required URL ...
}
