package auth

import "github.com/gin-gonic/gin"

type AuthController interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	UploadPhoto(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}
