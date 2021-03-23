package hotels

import (
	"gorm.io/gorm"
	"hotel_management/bookings"
)

type Hotel struct {
	*gorm.Model
	DisplayName string
	Description string
	Price float64
	Bookings []bookings.Booking
}