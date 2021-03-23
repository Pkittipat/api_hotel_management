package database

import (
	"fmt"
	"log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase(user, password, host, port, dbname string) {
	connectionString := 
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", host, user, password, dbname, port)
	
	database, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err);
	} else {
		fmt.Println("Connection Database...")
	}

	DB = database
}