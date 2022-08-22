package article

import (
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/rafiseptian90/GoArticle/app/models"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	"github.com/rafiseptian90/GoArticle/config"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"strconv"
	"strings"
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
	var articles map[string]interface{}

	if len(ctx.QueryArray("tags")) < 1 {
		articles = controller.repository.GetArticles()
	} else {
		//articles = controller.repository.GetArticlesByTags(ctx.QueryArray("tags"))
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
	authUser := models.AuthUser(ctx)
	var articleRequest models.ArticleRequest

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

func (controller *Controller) UploadThumbnail(ctx *gin.Context) {
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

func (controller *Controller) Update(ctx *gin.Context) {
	authUser := models.AuthUser(ctx)
	articleID, _ := strconv.Atoi(ctx.Param("articleID"))
	var articleRequest models.ArticleRequest

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

func (controller *Controller) Delete(ctx *gin.Context) {
	articleID, _ := strconv.Atoi(ctx.Param("articleID"))

	if err := controller.repository.DeleteArticle(articleID); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Article has been deleted")
	return
}
