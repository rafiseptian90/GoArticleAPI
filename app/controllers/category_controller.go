package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/handlers/requests"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"strconv"
)

type CategoryControllerInterface interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type CategoryController struct {
	repository *repositories.CategoryRepository
}

func NewCategoryController(repository *repositories.CategoryRepository) *CategoryController {
	return &CategoryController{
		repository: repository,
	}
}

func (controller *CategoryController) Index(ctx *gin.Context) {
	categories := controller.repository.GetCategories()

	ResponseJSON.SuccessWithData(ctx, "Categories has been loaded", categories)
}

func (controller *CategoryController) Show(ctx *gin.Context) {
	categoryID, _ := strconv.Atoi(ctx.Param("categoryID"))

	category, err := controller.repository.GetCategory(categoryID)
	if err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "Category has been loaded", category)
}

func (controller *CategoryController) Store(ctx *gin.Context) {
	var categoryRequest requests.CategoryRequest

	err := ctx.ShouldBindJSON(&categoryRequest)
	if err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	err = controller.repository.StoreCategory(&categoryRequest)
	if err != nil {
		ResponseJSON.InternalServerError(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "New Category has been added")
}

func (controller *CategoryController) Update(ctx *gin.Context) {
	var categoryRequest requests.CategoryRequest
	categoryID, _ := strconv.Atoi(ctx.Param("categoryID"))

	if err := ctx.ShouldBindJSON(&categoryRequest); err != nil {
		ResponseJSON.BadRequest(ctx, err.Error())
		return
	}

	if err := controller.repository.UpdateCategory(categoryID, &categoryRequest); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Category has been updated")
}

func (controller *CategoryController) Delete(ctx *gin.Context) {
	categoryID, _ := strconv.Atoi(ctx.Param("categoryID"))

	if err := controller.repository.DeleteCategory(categoryID); err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.Success(ctx, "Category has been deleted")
}
