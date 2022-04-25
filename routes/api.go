package routes

import (
	"github.com/gin-gonic/gin"
	"hanya-gin/app/controllers"
	"hanya-gin/app/middleware"
	"hanya-gin/app/services"
	"net/http"
)

func SetApiGroupRoutes(router *gin.RouterGroup) {

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	router.POST("/user/register", controllers.Register)
	router.POST("/user/login", controllers.Login)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/user/info", controllers.Info)
		authRouter.POST("/user/logout", controllers.Logout)
		authRouter.POST("/upload/image", controllers.UploadImage)
	}
}
