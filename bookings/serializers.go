package bookings

import (
	"hotel_management/utils"
)

type CreateBookingSerializer struct {
	HotelID uint `json:"hotel_id"`
}

type ResponseBooking struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	HotelID uint `json:"hotel_id"`
}


func (self *CreateBookingSerializer) Validate() map[string]interface{} {
	if self.HotelID == 0 {
		return utils.HandleResponse("Hotel Id field is required", 400)
	}
	return nil
}