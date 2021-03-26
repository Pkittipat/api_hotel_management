package bookings

import (
	"testing"
	"hotel_management/utils"
)

func TestValidateErrorNoDataRequest(t *testing.T) {
	req := CreateBookingSerializer{HotelID: 0}
	if req.Validate() == nil {
		t.Errorf("Validate() failed, expected %v, got %v", utils.HandleResponse("Hotel Id field is required", 400), req.Validate())
	} else {
		t.Logf("Validate() success")
	}
}

func TestDataRequestIsValid(t *testing.T) {
	req := CreateBookingSerializer{HotelID: 1}
	if req.Validate() != nil {
		t.Errorf("Validate() failed, expexted %v, got %v", map[string]interface{}{}, req.Validate())
	} else {
		t.Logf("Validate() success")
	}
}