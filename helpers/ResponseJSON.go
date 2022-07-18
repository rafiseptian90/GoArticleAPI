package ResponseJSON

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 200 /*
func Success(ctx *gin.Context, msgStatus string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code_status": http.StatusOK,
		"msg_status":  msgStatus,
	})
}

// SuccessWithData 200 /*
func SuccessWithData(ctx *gin.Context, msgStatus string, data any) {
	ctx.JSON(http.StatusOK, gin.H{
		"code_status": http.StatusOK,
		"msg_status":  msgStatus,
		"data":        data,
	})
}

// BadRequest 400 /*
func BadRequest(ctx *gin.Context, msgStatus string) {
	ctx.JSON(http.StatusBadRequest, gin.H{
		"code_status": http.StatusBadRequest,
		"msg_status":  msgStatus,
	})
}

// Unauthorized 401 /*
func Unauthorized(ctx *gin.Context, msgStatus string) {
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"code_status": http.StatusUnauthorized,
		"msg_status":  msgStatus,
	})
}

// Forbidden 403 /*
func Forbidden(ctx *gin.Context, msgStatus string) {
	ctx.JSON(http.StatusForbidden, gin.H{
		"code_status": http.StatusForbidden,
		"msg_status":  msgStatus,
	})
}

// NotFound 404 /*
func NotFound(ctx *gin.Context, msgStatus string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"code_status": http.StatusNotFound,
		"msg_status":  msgStatus,
	})
}

// InternalServerError 404 /*
func InternalServerError(ctx *gin.Context, msgStatus string) {
	ctx.JSON(http.StatusInternalServerError, gin.H{
		"code_status": http.StatusInternalServerError,
		"msg_status":  msgStatus,
	})
}
