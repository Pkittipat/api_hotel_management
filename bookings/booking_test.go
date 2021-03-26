package bookings

import (
	"testing"
)

func TestValidateErrorNoDataRequest(t *testing.T) {
	req := CreateBookingSerializer{}
	if req.Validate() == nil {
		t.Errorf("Validate() fail")
	} else {
		t.Logf("Validate() success")
	}
}

func TestDataRequestIsValid(t *testing.T) {
	req := CreateBookingSerializer{
		HotelID: 1,
	}
	if req.Validate() != nil {
		t.Errorf("Validate() fail")
	} else {
		t.Logf("Validate() success")
	}
}