package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/zenpk/dorm-system/internal/middleware"
	"github.com/zenpk/dorm-system/internal/service"
)

func InitRouter(router *gin.Engine) {
	router.Static("/static", "./assets/public") // 静态资源路径
	router.Use(middleware.CORSFilter())         // CORS

	// 不做任何登录验证的 URL
	router.GET("/importData/stutable", handler.studentInfo.GetAll)
	router.GET("/importData/stuquery", handler.studentInfo.GetAllByParams)
	router.GET("/importData/getastu", handler.studentInfo.GetById)
	router.POST("/importData/addstu", handler.studentInfo.Add)
	router.PUT("/importData/editstu", handler.studentInfo.Edit)
	router.DELETE("/importData/delstu", handler.studentInfo.Delete)
	router.POST("/login", handler.userCredential.Login)
	router.POST("/register", handler.userCredential.Register)
	router.GET("/user/logout", handler.userCredential.Logout)
	router.GET("/user/state", handler.userCredential.LoginState)

	// 不需要登录，但可以从 token 中获取用户信息的 URL
	routerNoAuth := router.Group("/")
	routerNoAuth.Use(middleware.CheckAuthInfo())
	{
		routerNoAuth.GET("/index", service.HelloDemo)
	}

	// 需要登录的 URL
	routerAuth := router.Group("/")
	routerAuth.Use(middleware.RequireLogin())
	{
		routerAuth.GET("/user-page", service.UserDemo)
		routerAuth.GET("/admin-page", service.AdminDemo)
		routerAuth.GET("/all-users", handler.userInfo.GetAllUserInfos)
		routerAuth.GET("/my-info", handler.userInfo.GetUserInfo)
	}
}
