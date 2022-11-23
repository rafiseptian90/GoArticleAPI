package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	v1 "github.com/rafiseptian90/GoArticle/api/v1"
	"github.com/rafiseptian90/GoArticle/pkg/config"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	DB := config.DBConnection()
	//database.InitMigration(DB)
	//database.InitSeeder(DB)

	app := gin.Default()
	app.Use(cors.Default())

	v1.NewAPIHandlerV1(app, DB)

	err := app.Run(":5050")
	if err != nil {
		log.Fatal(err.Error())
	}
}
