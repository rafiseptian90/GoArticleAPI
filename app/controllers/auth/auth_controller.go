package auth

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	Logout(ctx *gin.Context)
}
