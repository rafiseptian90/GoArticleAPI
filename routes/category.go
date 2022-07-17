package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafiseptian90/GoArticle/app/controllers"
	"github.com/rafiseptian90/GoArticle/app/repositories"
	"github.com/rafiseptian90/GoArticle/config"
)

func InitCategoryRoutes(router *gin.Engine) {

	DB := config.DBConnection()

	categoryRepository := repositories.CategoryRepository{DB: DB}
	categoryController := controllers.NewCategoryController(&categoryRepository)

	router.GET("/category", categoryController.Index)
	router.GET("/category/:categoryID", categoryController.Show)
	router.POST("/category", categoryController.Store)
	router.PUT("/category/:categoryID", categoryController.Update)
	router.DELETE("/category/:categoryID", categoryController.Delete)
}
