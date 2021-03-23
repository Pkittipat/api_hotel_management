package bookings

type CreateBookingSerializer struct {
	HotelID uint
}

type ResponseBooking struct {
	ID uint
	UserID uint
	HotelID uint
}