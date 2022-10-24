package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/middleware"
)

func InitRouter(router *gin.Engine) {
	//router.Static("/static", "./assets/public") // static resources
	router.Use(middleware.CORSFilter()) // CORS

	// no other middleware URL
	//router.POST("/login", ginHandler.userCredential.Login)
	router.POST("/register", ginHandler.userCredential.Register)
	//router.GET("/available-buildings", ginHandler.building.GetAvailableBuildings)
	//router.GET("/all-available-count", ginHandler.building.GetAvailableCount)

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
		//router.PATCH("/change-password", ginHandler.userCredential.UpdatePassword)
		//router.GET("/logout", ginHandler.userCredential.Logout)
	}
}
