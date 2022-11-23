package controller

import (
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	_articleRepository "github.com/rafiseptian90/GoArticle/api/v1/article/repository"
	models2 "github.com/rafiseptian90/GoArticle/models"
	"github.com/rafiseptian90/GoArticle/pkg/config"
	"github.com/rafiseptian90/GoArticle/pkg/helpers"
	"strconv"
	"strings"
)

type ArticleController struct {
	repository *_articleRepository.ArticleRepository
}

func NewArticleController(repository *_articleRepository.ArticleRepository) *ArticleController {
	return &ArticleController{
		repository: repository,
	}
}

func (controller *ArticleController) Index(ctx *gin.Context) {

	if len(ctx.QueryArray("tags")) < 1 {
		var articles map[string]interface{}

		articles = controller.repository.GetArticles()

		ResponseJSON.SuccessWithData(ctx, "Articles has been loaded", articles)
	} else {
		var articles []models2.Article

		switch ctx.Query("sortBy") {
		case "trending":
			articles = controller.repository.GetTrendingArticlesByTags(ctx.QueryArray("tags"))
			break
		case "latest":
			articles = controller.repository.GetLatestArticlesByTags(ctx.QueryArray("tags"))
			break
		case "best":
			articles = controller.repository.GetBestArticlesByTags(ctx.QueryArray("tags"))
			break
		default:
			articles = controller.repository.GetArticlesByTags(ctx.QueryArray("tags"))
		}

		ResponseJSON.SuccessWithData(ctx, "Articles has been loaded", articles)
	}
}

func (controller *ArticleController) Show(ctx *gin.Context) {
	articleSlug := ctx.Param("articleSlug")

	article, err := controller.repository.GetArticle(articleSlug)
	if err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Article has been loaded", article)
}

func (controller *ArticleController) Store(ctx *gin.Context) {
	authUser := models2.AuthUser(ctx)
	var articleRequest models2.ArticleRequest

	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	articleRequest.UserId = authUser.Id
	articleRequest.Slug = slug.Make(articleRequest.Title)

	if err := controller.repository.StoreArticle(&articleRequest); err != nil {
		ResponseJSON.InternalServerError(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "New Article has been added")
}

func (controller *ArticleController) UploadThumbnail(ctx *gin.Context) {
	// Init Cloudinary
	cld, err := config.InitCLD()
	if err != nil {
		ResponseJSON.InternalServerError(ctx, err.Error())
		return
	}

	name := strings.ReplaceAll(strings.ToLower(ctx.PostForm("title")), " ", "_")
	file, _, err := ctx.Request.FormFile("thumbnail")
	if err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	uploadResult, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: "articles/" + name})
	if err != nil {
		ResponseJSON.InternalServerError(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Article thumbnail has been uploaded", uploadResult.SecureURL)
	return
}

func (controller *ArticleController) Update(ctx *gin.Context) {
	authUser := models2.AuthUser(ctx)
	articleID, _ := strconv.Atoi(ctx.Param("articleID"))
	var articleRequest models2.ArticleRequest

	if err := ctx.ShouldBindJSON(&articleRequest); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	articleRequest.UserId = authUser.Id
	articleRequest.Slug = slug.Make(articleRequest.Title)

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
	return
}
