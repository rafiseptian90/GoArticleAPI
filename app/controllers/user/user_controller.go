package user

import "github.com/gin-gonic/gin"

type UserController interface {
	Show(ctx *gin.Context)
}
