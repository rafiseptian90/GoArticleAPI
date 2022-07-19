package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/controllers"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	"github.com/rafiseptian90/GoArticle/config"
)

func InitArticleRoutes(router *gin.Engine) {
	DB := config.DBConnection()
	articleRepository := repositories.NewArticleRepository(DB)
	articleController := controllers.NewArticleController(articleRepository)

	router.GET("/article", articleController.Index)
	router.GET("/article/:articleID", articleController.Show)
	router.POST("/article", articleController.Store)
	router.PUT("/article/:articleID", articleController.Update)
	router.DELETE("/article/:articleID", articleController.Delete)
}
