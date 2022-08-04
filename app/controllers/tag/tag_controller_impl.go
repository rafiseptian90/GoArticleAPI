package tag

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/models"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"strconv"
)

type Controller struct {
	repository *repositories.TagRepository
}

func NewTagController(repository *repositories.TagRepository) *Controller {
	return &Controller{
		repository: repository,
	}
}

func (controller *Controller) Index(ctx *gin.Context) {
	tags := controller.repository.GetTags()

	ResponseJSON.SuccessWithData(ctx, "Tags has been loaded", tags)
}

func (controller *Controller) Show(ctx *gin.Context) {
	tagID, _ := strconv.Atoi(ctx.Param("tagID"))

	tag, err := controller.repository.GetTag(tagID)
	if err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Tag has been loaded", tag)
}

func (controller *Controller) Store(ctx *gin.Context) {
	var tagRequest models.Tag

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

func (controller *Controller) Update(ctx *gin.Context) {
	var tagRequest models.Tag
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

func (controller *Controller) Delete(ctx *gin.Context) {
	tagID, _ := strconv.Atoi(ctx.Param("tagID"))

	if err := controller.repository.DeleteTag(tagID); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Tag has been deleted")
}
