package users

import (
	"time"
	"gorm.io/gorm"
	"hotel_management/bookings"
)

type User struct {
	*gorm.Model
	Username string	`json:"username"`
	Password string	`json:"password"`
	Email *string	`json:"email"`
	Account Account	`json:"account"`
	Bookings []bookings.Booking	`json:"bookings"`
}

type Account struct {
	*gorm.Model
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	BirthDate *time.Time `json:"birth_date"`
	UserID uint `json:"user_id"`
}