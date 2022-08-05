package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/controllers/auth"
	"github.com/rafiseptian90/GoArticle/config"
)

func InitAuthRoutes(router *gin.Engine) {
	DB := config.DBConnection()
	authController := auth.NewAuthController(DB)

	router.POST("/login", authController.Login)
	router.POST("/auth/register", authController.Register)
	router.POST("/forgot-password", authController.ForgotPassword)
	router.POST("/logout", authController.Logout)
}
