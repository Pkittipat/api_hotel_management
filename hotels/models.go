package hotels

import (
	"gorm.io/gorm"
	"hotel_management/bookings"
)

type Location struct {
	*gorm.Model
	Latitude *float64
	Longitude *float64
	Hotel Hotel
}

type Hotel struct {
	*gorm.Model
	DisplayName string
	Description string
	Price float64
	LocationID *uint
	Bookings []bookings.Booking
}