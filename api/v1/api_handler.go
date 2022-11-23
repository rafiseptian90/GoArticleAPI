package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/api/v1/article"
	_articleController "github.com/rafiseptian90/GoArticle/api/v1/article/controller"
	_articleRepository "github.com/rafiseptian90/GoArticle/api/v1/article/repository"
	"github.com/rafiseptian90/GoArticle/api/v1/tag"
	_tagController "github.com/rafiseptian90/GoArticle/api/v1/tag/controller"
	_tagRepository "github.com/rafiseptian90/GoArticle/api/v1/tag/repository"
	"github.com/rafiseptian90/GoArticle/api/v1/user"
	_userController "github.com/rafiseptian90/GoArticle/api/v1/user/controller"
	_userRepository "github.com/rafiseptian90/GoArticle/api/v1/user/repository"
	"gorm.io/gorm"
)

func NewAPIHandlerV1(router *gin.Engine, DB *gorm.DB) {
	// Tag Repository & Controller
	tagRepository := _tagRepository.NewTagRepository(DB)
	tagController := _tagController.NewTagController(tagRepository)

	// Article Repository & Controller
	articleRepository := _articleRepository.NewArticleRepository(DB)
	articleController := _articleController.NewArticleController(articleRepository)

	// User Repository & Controllers
	userRepository := _userRepository.NewUserRepository(DB)
	authController := _userController.NewAuthController(DB)
	userController := _userController.NewUserController(userRepository)

	apiRouter := router.Group("/api/v1")
	{
		tag.NewTagHandlers(apiRouter, tagController)
		article.NewArticleHandlers(apiRouter, articleController)
		user.NewUserHandlers(apiRouter, authController, userController)
	}
}
