package users

import (
	"time"
	"hotel_management/hotels"
	"hotel_management/utils"
)

type ResponseUser struct {
	ID uint `json:"id"`
	Username string	`json:"username"`
	Email *string `json:"email"`
	Account ResponseAccount `json:"account"`
}

type ResponseAccount struct {
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
}

type RequestAccount struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	BirthDate *time.Time `json:"birth_date,omitempty"`
}

type AuthenticateUserSerializer struct {
	Username string `json:"username"`
	Password string	`json:"password"`
}

type SignupUserSerializer struct {
	Username string `json:"username"`
	Password string	`json:"password"`
	ConfirmPassword string `json:"confirm_password"`
	Email *string `json:"email"`
	Account RequestAccount `json:"account"`
}


type ResponseBooking struct {
	ID uint `json:"id"`
	User ResponseUser `json:"user"`
	Hotel hotels.ResponseHotel `json:"hotel"`
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

	if self.ConfirmPassword == "" {
		return utils.HandleResponse("Confirm password field is required", 400)

	} else if self.Email == nil {
		return utils.HandleResponse("Email field is required", 400)
		
	} else if self.Account.FirstName == "" {
		return utils.HandleResponse("Firstname field is required", 400)

	} else if self.Account.LastName == "" {
		return utils.HandleResponse("Lastname field is required", 400)

	} else if self.Password != self.ConfirmPassword {
		return utils.HandleResponse("Password does not match", 400)
	}

	return nil
}