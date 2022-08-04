package tag

import (
	"github.com/gin-gonic/gin"
)

type TagController interface {
	Index(ctx *gin.Context)
	Show(ctx *gin.Context)
	Store(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
