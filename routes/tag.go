package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/controllers/tag"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	"github.com/rafiseptian90/GoArticle/config"
)

func InitTagRoutes(router *gin.Engine) {

	DB := config.DBConnection()

	tagRepository := repositories.TagRepository{DB: DB}
	tagController := tag.NewTagController(&tagRepository)

	router.GET("/tag", tagController.Index)
	router.GET("/tag/:tagID", tagController.Show)
	router.POST("/tag", tagController.Store)
	router.PUT("/tag/:tagID", tagController.Update)
	router.DELETE("/tag/:tagID", tagController.Delete)
}
