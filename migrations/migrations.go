package migrations

import (
	"hotel_management/users"
	"hotel_management/hotels"
	"hotel_management/database"
)

func createMockHotels() {
	hotels := &[2]hotels.Hotel{
		{DisplayName: "The Eight Serviced Apartment", Description: "Mahachai, Mueang Samut Sakhon District, Samut Sakhon 74000•092 470 2888", Price: 851},
		{DisplayName: "Thongchen Residence", Description: "Contemporary serviced apartments with kitchenettes, flat-screens & free Wi-Fi, as well as parking.", Price: 850},
	}
	database.DB.Create(&hotels)
}

func Migrate() {
	user := users.User{}
	account := users.Account{}
	hotel := hotels.Hotel{}
	
	database.DB.AutoMigrate(&user, &account, &hotel)
	// createMockHotels()
}