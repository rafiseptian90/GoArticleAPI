package user

import (
	"github.com/gin-gonic/gin"
	controller2 "github.com/rafiseptian90/GoArticle/api/v1/user/controller"
	"github.com/rafiseptian90/GoArticle/api/v1/user/middleware"
)

func NewUserHandlers(router *gin.RouterGroup, authController *controller2.AuthController, userController *controller2.UserController) {
	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/register", authController.Register)
		authRouter.POST("/forgot-password", middleware.JWTAuthMiddleware(), authController.ForgotPassword)
		authRouter.POST("/refresh", middleware.JWTAuthMiddleware(), authController.Refresh)
	}

	router.GET("/user/:username", userController.FindByUsername)
	router.PUT("/user/profile", middleware.JWTAuthMiddleware(), authController.UpdateProfile)
	router.POST("/user/profile/upload", middleware.JWTAuthMiddleware(), authController.UploadPhoto)
}
