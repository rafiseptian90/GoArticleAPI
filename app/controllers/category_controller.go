package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/repositories"
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
	//TODO implement me
	panic("implement me")
}

func (controller *CategoryController) Show(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *CategoryController) Store(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *CategoryController) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (controller *CategoryController) Delete(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
