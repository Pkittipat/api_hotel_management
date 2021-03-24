package hotels

import (
	"gorm.io/gorm"
	"hotel_management/bookings"
)

type Location struct {
	*gorm.Model
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Hotel Hotel `json:"hotel"`
}

type Detail struct {
	*gorm.Model
	Title string `json:"title"`
}

type Hotel struct {
	*gorm.Model
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	LocationID *uint `json:"location_id"`
	Bookings []bookings.Booking `json:"bookings"`
	Details []Detail `json:"detials" gorm:"many2many:hotel_detail"`
}