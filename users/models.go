package users

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	*gorm.Model
	Username string
	Password string
	Email *string
	Account Account
}

type Account struct {
	*gorm.Model
	FirstName string
	LastName string
	BirthDate *time.Time
	UserID uint
}