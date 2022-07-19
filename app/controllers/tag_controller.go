package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/handlers/requests"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"strconv"
)

type TagControllerInterface interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type TagController struct {
	repository *repositories.TagRepository
}

func NewTagController(repository *repositories.TagRepository) *TagController {
	return &TagController{
		repository: repository,
	}
}

func (controller *TagController) Index(ctx *gin.Context) {
	tags := controller.repository.GetTags()

	ResponseJSON.SuccessWithData(ctx, "Tags has been loaded", tags)
}

func (controller *TagController) Show(ctx *gin.Context) {
	tagID, _ := strconv.Atoi(ctx.Param("tagID"))

	tag, err := controller.repository.GetTag(tagID)
	if err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Tag has been loaded", tag)
}

func (controller *TagController) Store(ctx *gin.Context) {
	var tagRequest requests.TagRequest

	err := ctx.ShouldBindJSON(&tagRequest)
	if err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	err = controller.repository.StoreTag(&tagRequest)
	if err != nil {
		ResponseJSON.InternalServerError(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "New Tag has been added")
}

func (controller *TagController) Update(ctx *gin.Context) {
	var tagRequest requests.TagRequest
	tagID, _ := strconv.Atoi(ctx.Param("tagID"))

	if err := ctx.ShouldBindJSON(&tagRequest); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	if err := controller.repository.UpdateTag(tagID, &tagRequest); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Tag has been updated")
}

func (controller *TagController) Delete(ctx *gin.Context) {
	tagID, _ := strconv.Atoi(ctx.Param("tagID"))

	if err := controller.repository.DeleteTag(tagID); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Tag has been deleted")
}
