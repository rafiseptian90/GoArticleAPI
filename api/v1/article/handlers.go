package article

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/api/v1/article/controller"
	"github.com/rafiseptian90/GoArticle/api/v1/user/middleware"
)

func NewArticleHandlers(router *gin.RouterGroup, articleController *controller.ArticleController) {
	router.GET("/article", articleController.Index)
	router.GET("/article/:articleSlug", articleController.Show)
	router.POST("/article/thumbnail/upload", middleware.JWTAuthMiddleware(), articleController.UploadThumbnail)
	router.POST("/article", middleware.JWTAuthMiddleware(), articleController.Store)
	router.PUT("/article/:articleID", middleware.JWTAuthMiddleware(), articleController.Update)
	router.DELETE("/article/:articleID", middleware.JWTAuthMiddleware(), articleController.Delete)
}
