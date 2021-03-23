package users

import (
	"time"
	"gorm.io/gorm"
	"hotel_management/bookings"
)

type User struct {
	*gorm.Model
	Username string
	Password string
	Email *string
	Account Account
	Bookings []bookings.Booking
}

type Account struct {
	*gorm.Model
	FirstName string
	LastName string
	BirthDate *time.Time
	UserID uint
}