package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/controllers/auth"
	"github.com/rafiseptian90/GoArticle/config"
)

func InitAuthRoutes(router *gin.Engine) {
	DB := config.DBConnection()
	authController := auth.NewAuthController(DB)

	authRouter := router.Group("/auth")
	{
		authRouter.POST("/login", authController.Login)
		authRouter.POST("/register", authController.Register)
		authRouter.POST("/forgot-password", authController.ForgotPassword)
		authRouter.POST("/logout", authController.Logout)
	}
}
