package tag

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"github.com/rafiseptian90/GoArticle/app/models"
	"github.com/rafiseptian90/GoArticle/app/repositories/tag"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
)

type Controller struct {
	repository *tag.Repository
}

func NewTagController(repository *tag.Repository) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (controller *Controller) Index(ctx *gin.Context) {
	tags := controller.repository.GetTags()

	ResponseJSON.SuccessWithData(ctx, "Tags has been loaded", tags)
}

func (controller *Controller) Show(ctx *gin.Context) {
	tag, err := controller.repository.GetTag(ctx.Param("tagSlug"))
	if err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Tag has been loaded", tag)
}

func (controller *Controller) Store(ctx *gin.Context) {
	var tagRequest models.TagRequest

	err := ctx.ShouldBindJSON(&tagRequest)
	if err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	tagRequest.Slug = slug.Make(tagRequest.Name)

	err = controller.repository.StoreTag(&tagRequest)
	if err != nil {
		ResponseJSON.InternalServerError(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "New Tag has been added")
}

func (controller *Controller) Update(ctx *gin.Context) {
	var tagRequest models.TagRequest

	if err := ctx.ShouldBindJSON(&tagRequest); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	tagRequest.Slug = slug.Make(tagRequest.Name)

	if err := controller.repository.UpdateTag(ctx.Param("tagSlug"), &tagRequest); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Tag has been updated")
}

func (controller *Controller) Delete(ctx *gin.Context) {
	if err := controller.repository.DeleteTag(ctx.Param("tagSlug")); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Tag has been deleted")
}
