package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rafiseptian90/GoArticle/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Uncomment this three line below if you want to run a migration and seeders
	//DB := config.DBConnection()
	//database.InitMigration(DB)
	//database.InitSeeder(DB)

	router := gin.Default()
	router.Use(cors.Default())
	router.Static("/public", "./public")
	routes.InitRoutes(router)

	err = router.Run(":5050")
	if err != nil {
		log.Fatal(err.Error())
	}
}
