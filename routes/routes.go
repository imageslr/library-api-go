package routes

import (
	"library-api/config"
	"library-api/controllers/test"
	"library-api/controllers/user"

	"github.com/gin-gonic/gin"
)

// Register 注册路由
func Register(router *gin.Engine) {
	apiPrefix := config.AppConfig.APIPrefix
	api := router.Group(apiPrefix)
	{
		api.GET("/ping", test.Ping)

		api.POST("/codes", user.SendCode)
		api.POST("/login", user.Login)
	}
}
