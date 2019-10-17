package routes

import (
	"library-api/config"
	"library-api/controllers/image"
	"library-api/controllers/test"
	"library-api/controllers/user"
	"library-api/middlewares"

	"github.com/gin-gonic/gin"
)

// Register 注册路由
func Register(router *gin.Engine) {
	apiPrefix := config.AppConfig.APIPrefix

	// 静态文件
	router.Static("/upload", "storage/upload")

	api := router.Group(apiPrefix)
	{
		api.GET("/ping", test.Ping)

		// 注册登录
		api.POST("/codes", user.SendCode)
		api.POST("/login", user.Login)
	}

	authAPI := router.Group(apiPrefix, middlewares.Auth)
	{
		// 用户信息
		authAPI.GET("/user", user.CurrentUser)
		authAPI.POST("/user", user.UpdateCurrentUser)

		// 上传图片
		authAPI.POST("/upload/image", image.Upload)
	}

}
