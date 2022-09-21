package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/controllers/article"
	"github.com/rafiseptian90/GoArticle/app/middleware"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	"github.com/rafiseptian90/GoArticle/config"
)

func InitArticleRoutes(router *gin.RouterGroup) {
	DB := config.DBConnection()
	articleRepository := repositories.NewArticleRepository(DB)
	articleController := article.NewArticleController(articleRepository)

	router.GET("/article", articleController.Index)
	router.GET("/article/:articleSlug", articleController.Show)
	router.POST("/article/thumbnail/upload", middleware.JWTAuthMiddleware(), articleController.UploadThumbnail)
	router.POST("/article", middleware.JWTAuthMiddleware(), articleController.Store)
	router.PUT("/article/:articleID", middleware.JWTAuthMiddleware(), articleController.Update)
	router.DELETE("/article/:articleID", middleware.JWTAuthMiddleware(), articleController.Delete)
}
