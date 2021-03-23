package bookings

import (
	"gorm.io/gorm"
)

type Booking struct {
	*gorm.Model
	UserID uint `json:"user_id"`
	HotelID uint `json:"hotel_id"`
}