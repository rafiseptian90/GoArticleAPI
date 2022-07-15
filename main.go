package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafiseptian90/GoArticle/config"
	"github.com/rafiseptian90/GoArticle/database"
	"github.com/rafiseptian90/GoArticle/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := config.DBConnection()
	database.InitMigration(DB)

	router := gin.Default()
	routes.InitRoutes(router)

	err = router.Run(":8000")
	if err != nil {
		log.Fatal(err.Error())
	}
}
