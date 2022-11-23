package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rafiseptian90/GoArticle/app/controllers/user"
	repository "github.com/rafiseptian90/GoArticle/app/repositories/user"
	"github.com/rafiseptian90/GoArticle/config"
)

func InitUserRoutes(router *gin.RouterGroup) {
	DB := config.DBConnection()

	userRepository := repository.NewUserRepository(DB)
	userController := controller.NewUserController(userRepository)

	router.GET("/user/:username", userController.Show)
}
