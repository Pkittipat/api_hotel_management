package hotels

type ResponseLocation struct {
	Latitude *float64
	Longitude *float64
}

type ResponseHotel struct {
	ID	uint
	DisplayName string
	Description string
	Price float64
}

type ResponseInfoHotel struct {
	ID	uint
	DisplayName string
	Description string
	Price float64
	Location ResponseLocation
}