package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/controllers"
)

func InitCategoryRoutes(router *gin.Engine) {
	categoryController := controllers.NewCategoryController()

	router.GET("/category", categoryController.Index)
	router.GET("/category/:categoryID", categoryController.Show)
	router.POST("/category", categoryController.Store)
	router.PUT("/category/:categoryID", categoryController.Update)
	router.DELETE("/category/:categoryID", categoryController.Delete)
}
