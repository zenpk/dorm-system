package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/middleware"
)

func InitRouter(router *gin.Engine) {
	//router.Static("/static", "./assets/public") // static resources
	router.Use(middleware.CORSFilter()) // CORS

	// no middleware URL
	router.POST("/login", ginHandler.user.Login)
	router.POST("/register", ginHandler.user.Register)
	router.GET("/available", ginHandler.dorm.GetAvailableNum)

	// login required URL
	routerAuth := router.Group("/")
	routerAuth.Use(middleware.RequireLogin())
	{
		//router.GET("/my-info", ginHandler.userInfo.GetMyInfo)
		//router.PATCH("/change-password", ginHandler.user.UpdatePassword)
		router.GET("/logout", ginHandler.user.Logout)
		router.POST("/order-create", ginHandler.order.Submit)
		router.POST("/team-create", ginHandler.team.Create)
	}

	// admin required URL

}
