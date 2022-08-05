package routes

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	InitTagRoutes(router)
	InitArticleRoutes(router)
	InitAuthRoutes(router)
}
