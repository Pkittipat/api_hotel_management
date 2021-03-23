package bookings

type CreateBookingSerializer struct {
	HotelID uint `json:"hotel_id"`
}

type ResponseBooking struct {
	ID uint `json:"id"`
	UserID uint `json:"user_id"`
	HotelID uint `json:"hotel_id"`
}