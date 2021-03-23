package users

import (
	"time"
	"hotel_management/hotels"
	"hotel_management/utils"
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
	ConfirmPassword string
	Email *string
	Account RequestAccount
}


type ResponseBooking struct {
	ID uint
	User ResponseUser
	Hotel hotels.ResponseHotel
}

func validateUsernameAndPassword(username, password string) map[string]interface{} {
	if username == "" {
		return utils.HandleResponse("username field is required", 400)
	} else if password == "" {
		return utils.HandleResponse("password field is required", 400)
	}
	return nil
}

func (self *AuthenticateUserSerializer) Validate() map[string]interface{} {
	return validateUsernameAndPassword(self.Username, self.Password)
}


func (self *SignupUserSerializer) Validate() map[string]interface{} {
	err := validateUsernameAndPassword(self.Username, self.Password)
	if err != nil {
		return err
	}

	if self.Password != self.ConfirmPassword {
		return utils.HandleResponse("Password does not match", 400)
	}

	return nil
}