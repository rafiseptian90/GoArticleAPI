package routes

import "github.com/gin-gonic/gin"

func InitRoutes(router *gin.Engine) {
	apiRouter := router.Group("api/v1")
	{
		InitTagRoutes(apiRouter)
		InitArticleRoutes(apiRouter)
		InitAuthRoutes(apiRouter)
		InitUserRoutes(apiRouter)
	}
}
