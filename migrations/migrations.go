package migrations

import (
	"hotel_management/users"
	"hotel_management/hotels"
	"hotel_management/database"
	"hotel_management/bookings"
)


func createMockHotels() {
	
	locations := &[2]hotels.Location{
		{Latitude: 13.562505246598713, Longitude: 100.27970043047608},
		{Latitude: 13.551410693364241, Longitude: 100.27834312691078},
	}
	database.DB.Create(&locations)
	
	hotels := &[2]hotels.Hotel{
		{
			DisplayName: "The Eight Serviced Apartment", 
			Description: "Mahachai, Mueang Samut Sakhon District, Samut Sakhon 74000•092 470 2888", 
			Price: 851, 
			LocationID: &locations[0].ID,
			Details: []hotels.Detail{
				{Title: "Free WiFi"},
				{Title: "Safe"},
				{Title: "Shower"},
			},
		},
		{
			DisplayName: "Thongchen Residence", 
			Description: "Contemporary serviced apartments with kitchenettes, flat-screens & free Wi-Fi, as well as parking.", 
			Price: 850, 
			LocationID: &locations[1].ID,
			Details: []hotels.Detail{
				{Title: "Mini-fridge"},
				{Title: "Safe"},
				{Title: "Shower"},
				{Title: "Satellite LCD TV"},
			},
		},
	}
	database.DB.Create(&hotels)
	
}

func Migrate() {
	User := users.User{}
	Account := users.Account{}
	Hotel := hotels.Hotel{}
	Booking := bookings.Booking{}
	Location := hotels.Location{}
	Detail := hotels.Detail{}
	
	database.DB.AutoMigrate(&User, &Account, &Location, &Detail, &Hotel, &Booking)
	// createMockHotels()
}