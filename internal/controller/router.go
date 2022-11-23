package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/middleware"
)

func InitRouter(router *gin.Engine) {
	//router.Static("/static", "./assets/public") // static resources
	router.Use(middleware.CORSFilter()) // CORS
	router.GET("/temp", tempHandler)    // for testing

	// no middleware URL
	router.POST("/login", ginHandler.user.Login)
	router.POST("/register", ginHandler.user.Register)
	router.GET("/remain-cnt", ginHandler.dorm.GetRemainCnt) // get the number of remaining beds count in each building

	// login required URL
	routerAuth := router.Group("/")
	routerAuth.Use(middleware.RequireLogin())
	{
		//router.GET("/my-info", ginHandler.userInfo.GetMyInfo)
		//router.PATCH("/change-password", ginHandler.user.UpdatePassword)
		routerAuth.GET("/logout", ginHandler.user.Logout)
		routerAuth.POST("/team-create", ginHandler.team.Create)
		routerAuth.GET("/team", ginHandler.team.Get) // get one's team info
		routerAuth.POST("/team-join", ginHandler.team.Join)
		routerAuth.POST("/order-create", ginHandler.order.Submit)
		routerAuth.GET("/order", ginHandler.order.Get) // get one's team's order info
	}

	// admin required URL

}
