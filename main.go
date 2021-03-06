package main

import (
	"os"
	"log"
	"hotel_management/api"
	"hotel_management/database"
	"hotel_management/migrations"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/.env")
  	if err != nil {
    	log.Fatal("Error loading .env file")
  	}
	
	database.InitDatabase(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_NAME"));

	migrations.Migrate()
	api.Run()
}