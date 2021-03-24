package migrations

import (
	"hotel_management/users"
	"hotel_management/hotels"
	"hotel_management/database"
	"hotel_management/bookings"
)


func createMockHotels() {
	
	locations := &[4]hotels.Location{
		{Latitude: 13.562505246598713, Longitude: 100.27970043047608},
		{Latitude: 13.551410693364241, Longitude: 100.27834312691078},
		{Latitude: 13.7072377, Longitude: 100.3766593},
		{Latitude: 13.6939444, Longitude: 100.3365937},
	}
	database.DB.Create(&locations)
	
	hotels := &[4]hotels.Hotel{
		{
			DisplayName: "The Eight Serviced Apartment", 
			Description: "Mahachai, Mueang Samut Sakhon District, Samut Sakhon 74000•092 470 2888", 
			Price: 851, 
			LocationID: &locations[0].ID,
			Details: []hotels.Detail{
				{Title: "Mini-fridge"},
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
				{Title: "Free WiFi"},
				{Title: "Safe"},
				{Title: "Shower"},
				{Title: "Satellite LCD TV"},
			},
		},
		{
			DisplayName: "Chill Apartment", 
			Description: "Chill Apartment, 46 Phet Kasem 67 Alley, Lane 4, Lak Song, Bang Khae, Bangkok 10160", 
			Price: 480, 
			LocationID: &locations[2].ID,
			Details: []hotels.Detail{
				{Title: "Free WiFi"},
				{Title: "Safe"},
				{Title: "Shower"},
				{Title: "Satellite LCD TV"},
			},
		},
		{
			DisplayName: "P29", 
			Description: "55/7 Liap Khlong Phasi Charoen Fang Nuea 10 Alley, Nong Khaem, Bangkok 10160", 
			Price: 450, 
			LocationID: &locations[3].ID,
			Details: []hotels.Detail{
				{Title: "Free WiFi"},
				{Title: "Safe"},
				{Title: "Shower"},
				{Title: "Mini-fridge"},
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