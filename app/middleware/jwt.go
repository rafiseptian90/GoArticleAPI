package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/config"
	ResponseJSON "github.com/rafiseptian90/GoArticle/helpers"
	"strings"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		if token == "" {
			ResponseJSON.Unauthorized(ctx, "Authorization token is not found !")
			ctx.Abort()
			return
		}

		err := config.JWTValidateToken(token)
		if err != nil {
			ResponseJSON.Forbidden(ctx, err.Error())
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
