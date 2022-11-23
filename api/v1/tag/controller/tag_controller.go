package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	_tagRepository "github.com/rafiseptian90/GoArticle/api/v1/tag/repository"
	"github.com/rafiseptian90/GoArticle/models"
	"github.com/rafiseptian90/GoArticle/pkg/helpers"
)

type TagController struct {
	repository *_tagRepository.TagRepository
}

func NewTagController(repository *_tagRepository.TagRepository) *TagController {
	return &TagController{
		repository: repository,
	}
}

func (controller *TagController) Index(ctx *gin.Context) {
	tags := controller.repository.GetTags()

	ResponseJSON.SuccessWithData(ctx, "Tags has been loaded", tags)
}

func (controller *TagController) Show(ctx *gin.Context) {
	tag, err := controller.repository.GetTag(ctx.Param("tagSlug"))
	if err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Tag has been loaded", tag)
}

func (controller *TagController) Store(ctx *gin.Context) {
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

func (controller *TagController) Update(ctx *gin.Context) {
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

func (controller *TagController) Delete(ctx *gin.Context) {
	if err := controller.repository.DeleteTag(ctx.Param("tagSlug")); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Tag has been deleted")
}
