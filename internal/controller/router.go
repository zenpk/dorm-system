package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/middleware"
)

func InitRouter(router *gin.Engine) {
	//router.Static("/static", "./assets/public") // static resources
	router.Use(middleware.CORSFilter()) // CORS

	// no other middleware URL
	router.POST("/login", ginHandler.user.Login)
	router.POST("/register", ginHandler.user.Register)
	router.GET("/available", ginHandler.dorm.GetAvailableNum)

	// not login required but can extract information from token
	routerNoAuth := router.Group("/")
	routerNoAuth.Use(middleware.CheckAuthInfo())
	{
	}

	// login required URL
	routerAuth := router.Group("/")
	routerAuth.Use(middleware.RequireLogin())
	{
		//router.GET("/my-info", ginHandler.userInfo.GetMyInfo)
		//router.PATCH("/change-password", ginHandler.user.UpdatePassword)
		router.GET("/logout", ginHandler.user.Logout)
		router.POST("/order-create", ginHandler.order.Submit)
	}
}
