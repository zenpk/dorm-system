package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/middleware"
)

func InitRouter(router *gin.Engine) {
	//router.Static("/static", "./assets/public") // static resources
	router.Use(middleware.CORSFilter()) // CORS

	// no other middleware URL
	router.POST("/login", handler.userCredential.Login)
	router.POST("/register", handler.userCredential.Register)
	router.GET("/available-buildings", handler.building.GetAvailableBuildings)
	router.GET("/all-available-count", handler.building.GetAvailableCount)

	// not login required but can extract information from token
	routerNoAuth := router.Group("/")
	routerNoAuth.Use(middleware.CheckAuthInfo())
	{
	}

	// login required URL
	routerAuth := router.Group("/")
	routerAuth.Use(middleware.RequireLogin())
	{
		router.GET("/my-info", handler.userInfo.GetMyInfo)
		router.PATCH("/change-password", handler.userCredential.UpdatePassword)
		router.GET("/logout", handler.userCredential.Logout)
	}
}