package article

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/models"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"strconv"
)

type Controller struct {
	repository *repositories.ArticleRepository
}

func NewArticleController(repository *repositories.ArticleRepository) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (controller *Controller) Index(ctx *gin.Context) {
	var articles []models.Article

	if len(ctx.QueryArray("tags")) < 1 {
		articles = controller.repository.GetArticles()
	} else {
		articles = controller.repository.GetArticlesByTags(ctx.QueryArray("tags"))
	}

	ResponseJSON.SuccessWithData(ctx, "Articles has been loaded", articles)
}

func (controller *Controller) Show(ctx *gin.Context) {
	articleID, _ := strconv.Atoi(ctx.Param("articleID"))

	article, err := controller.repository.GetArticle(articleID)
	if err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Article has been loaded", article)
}

func (controller *Controller) Store(ctx *gin.Context) {
	var articleRequest models.Article

	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	if err := controller.repository.StoreArticle(&articleRequest); err != nil {
		ResponseJSON.InternalServerError(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "New Article has been added")
}

func (controller *Controller) Update(ctx *gin.Context) {
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

func (controller *Controller) Delete(ctx *gin.Context) {
	articleID, _ := strconv.Atoi(ctx.Param("articleID"))

	if err := controller.repository.DeleteArticle(articleID); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Article has been deleted")
}
