package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/controllers/auth"
	"github.com/rafiseptian90/GoArticle/app/middleware"
	"github.com/rafiseptian90/GoArticle/config"
)

func InitAuthRoutes(router *gin.RouterGroup) {
	DB := config.DBConnection()
	authController := auth.NewAuthController(DB)

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/register", authController.Register)
		authRouter.PUT("/update-profile", middleware.JWTAuthMiddleware(), authController.UpdateProfile)
		authRouter.POST("/update-profile/upload", middleware.JWTAuthMiddleware(), authController.UploadPhoto)
		authRouter.POST("/forgot-password", middleware.JWTAuthMiddleware(), authController.ForgotPassword)
		authRouter.POST("/refresh", middleware.JWTAuthMiddleware(), authController.Refresh)
	}
}
