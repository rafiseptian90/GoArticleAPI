package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rafiseptian90/GoArticle/config"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB := config.DBConnection()

	fmt.Println(DB)
}
