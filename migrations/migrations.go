package migrations

import (
	"hotel_management/users"
	"hotel_management/database"
)

func Migrate() {
	user := users.User{}
	account := users.Account{}
	
	database.DB.AutoMigrate(&user, &account)
}