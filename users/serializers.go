package users

import (
	"time"
	"hotel_management/hotels"
)

type ResponseUser struct {
	ID uint
	Username string
	Email *string
	Account ResponseAccount
}

type ResponseAccount struct {
	ID uint
	FirstName string
	LastName string
	BirthDate *time.Time
}

type RequestAccount struct {
	FirstName string
	LastName string
	BirthDate *time.Time
}

type AuthenticateUserSerializer struct {
	Username string
	Password string
}

type SignupUserSerializer struct {
	Username string
	Password string
	Email *string
	Account RequestAccount
}


type ResponseBooking struct {
	ID uint
	User ResponseUser
	Hotel hotels.ResponseHotel
}