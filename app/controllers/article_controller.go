package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/models"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"strconv"
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
	var articles []models.Article

	if len(ctx.QueryArray("tags")) < 1 {
		articles = controller.repository.GetArticles()
	} else {
		articles = controller.repository.GetArticlesByTags(ctx.QueryArray("tags"))
	}

	ResponseJSON.SuccessWithData(ctx, "Articles has been loaded", articles)
}

func (controller *ArticleController) Show(ctx *gin.Context) {
	articleID, _ := strconv.Atoi(ctx.Param("articleID"))

	article, err := controller.repository.GetArticle(articleID)
	if err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Article has been loaded", article)
}

func (controller *ArticleController) Store(ctx *gin.Context) {
	var articleRequest models.Article

	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	if err := controller.repository.StoreArticle(&articleRequest); err != nil {
		ResponseJSON.InternalServerError(ctx, err.Error())
	}

	ResponseJSON.Success(ctx, "New Article has been added")
}

func (controller *ArticleController) Update(ctx *gin.Context) {
	articleID, _ := strconv.Atoi(ctx.Param("articleID"))
	var articleRequest models.Article

	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	if err := controller.repository.UpdateArticle(articleID, &articleRequest); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Article has been updated")
}

func (controller *ArticleController) Delete(ctx *gin.Context) {
	articleID, _ := strconv.Atoi(ctx.Param("articleID"))

	if err := controller.repository.DeleteArticle(articleID); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Article has been deleted")
}
