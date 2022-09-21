package user

import (
	"github.com/gin-gonic/gin"
	userRepository "github.com/rafiseptian90/GoArticle/app/repositories/user"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"log"
)

type Controller struct {
	repository *userRepository.Repository
}

func NewUserController(repository *userRepository.Repository) *Controller {
	return &Controller{repository: repository}
}

func (controller *Controller) Show(ctx *gin.Context) {
	user, err := controller.repository.GetUserByUsername(ctx.Param("username"))
	if err != nil {
		ResponseJSON.NotFound(ctx, "Username is not found")
		log.Fatalf("Error :  %v", err)
		return
	}

	ResponseJSON.SuccessWithData(ctx, "User has been loaded", user)
	return
}
