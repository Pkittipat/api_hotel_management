package hotels

type ResponseLocation struct {
	Latitude *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}

type ResponseHotel struct {
	ID	uint `json:"id"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
}

type ResponseInfoHotel struct {
	ID	uint `json:"id"`
	DisplayName string `json:"display_name"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	Location ResponseLocation `json:"location"`
}