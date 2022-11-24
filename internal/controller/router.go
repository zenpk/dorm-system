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

	// login required URL
	routerAuth := router.Group("/")
	routerAuth.Use(middleware.RequireLogin())
	{
		routerAuth.GET("/logout", ginHandler.user.Logout)
		routerAuth.POST("/team-create", ginHandler.team.Create)
		routerAuth.GET("/team", ginHandler.team.Get) // get one's team info
		routerAuth.POST("/team-join", ginHandler.team.Join)
		routerAuth.POST("/order-create", ginHandler.order.Submit)
		routerAuth.GET("/order", ginHandler.order.Get) // get one's team's order info
	}

	// admin required URL

}
