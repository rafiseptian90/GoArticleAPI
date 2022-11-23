package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/rafiseptian90/GoArticle/app/controllers/tag"
	repository "github.com/rafiseptian90/GoArticle/app/repositories/tag"
	"github.com/rafiseptian90/GoArticle/config"
)

func InitTagRoutes(router *gin.RouterGroup) {

	DB := config.DBConnection()

	tagRepository := repository.NewTagRepository(DB)
	tagController := controller.NewTagController(tagRepository)

	router.GET("/tag", tagController.Index)
	router.GET("/tag/:tagSlug", tagController.Show)
	router.POST("/tag", tagController.Store)
	router.PUT("/tag/:tagSlug", tagController.Update)
	router.DELETE("/tag/:tagSlug", tagController.Delete)
}
