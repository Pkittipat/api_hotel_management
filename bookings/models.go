package bookings

import (
	"gorm.io/gorm"
)

type Booking struct {
	*gorm.Model
	UserID uint
	HotelID uint
}