package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/api/v1/user/repository"
	"github.com/rafiseptian90/GoArticle/pkg/helpers"
)

type UserController struct {
	repository *repository.UserRepository
}

func NewUserController(repository *repository.UserRepository) *UserController {
	return &UserController{repository: repository}
}

func (controller *UserController) FindByUsername(ctx *gin.Context) {
	user, err := controller.repository.FindByUsername(ctx.Param("username"))
	if err != nil {
		ResponseJSON.NotFound(ctx, err.Error())
		return
	}

	ResponseJSON.SuccessWithData(ctx, "User has been loaded", user)
	return
}
