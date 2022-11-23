package tag

import (
	"github.com/gin-gonic/gin"
	_tagController "github.com/rafiseptian90/GoArticle/api/v1/tag/controller"
)

func NewTagHandlers(router *gin.RouterGroup, tagController *_tagController.TagController) {
	router.GET("/tag", tagController.Index)
	router.GET("/tag/:tagSlug", tagController.Show)
	router.POST("/tag", tagController.Store)
	router.PUT("/tag/:tagSlug", tagController.Update)
	router.DELETE("/tag/:tagSlug", tagController.Delete)
}
