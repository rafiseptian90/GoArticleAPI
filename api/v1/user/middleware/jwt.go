package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/pkg/config"
	"github.com/rafiseptian90/GoArticle/pkg/helpers"
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

		userEmail, err := config.JWTValidateToken(token)
		if err != nil {
			ResponseJSON.Unauthorized(ctx, err.Error())
			ctx.Abort()
			return
		}

		ctx.Set("user.email", userEmail)
		ctx.Next()
	}
}
