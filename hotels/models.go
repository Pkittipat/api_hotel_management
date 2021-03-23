package hotels

import (
	"gorm.io/gorm"
)

type Hotel struct {
	*gorm.Model
	DisplayName string
	Description string
	Price float64
}