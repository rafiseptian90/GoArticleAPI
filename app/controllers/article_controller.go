package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/repositories"
)

type ArticleControllerInterface interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type ArticleController struct {
	repository *repositories.ArticleRepository
}

func NewArticleController(repository *repositories.ArticleRepository) *ArticleController {
	return &ArticleController{
		repository: repository,
	}
}

func (controller *ArticleController) Index(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *ArticleController) Show(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *ArticleController) Store(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *ArticleController) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *ArticleController) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
