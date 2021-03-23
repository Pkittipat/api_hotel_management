package api

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"hotel_management/users"
	"hotel_management/hotels"
	"hotel_management/bookings"
	"hotel_management/database"
	"hotel_management/utils"
	"encoding/json"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users/authenticate", AuthenticateUser).Methods("POST")
	router.HandleFunc("/api/users/signup", RegistrationUser).Methods("POST")
	router.HandleFunc("/api/users/me", GetUser).Methods("GET")
	router.HandleFunc("/api/users/bookings", GetListUserBookings).Methods("GET")
	router.HandleFunc("/api/hotels", GetListHotels).Methods("GET")
	router.HandleFunc("/api/hotels/{id}", GetInfoHotel).Methods("GET")
	router.HandleFunc("/api/bookings", CreateBooking).Methods("POST")
	fmt.Println("App is running on port :8000")
	log.Fatal(http.ListenAndServe(":8000", router));
}


func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	requestAuth := users.AuthenticateUserSerializer{}
	json.NewDecoder(r.Body).Decode(&requestAuth)

	response := users.AuthenticateUser(requestAuth.Username, requestAuth.Password)
	json.NewEncoder(w).Encode(response)
}

func RegistrationUser(w http.ResponseWriter, r *http.Request) {
	requestRegister := users.SignupUserSerializer{}
	json.NewDecoder(r.Body).Decode(&requestRegister)

	response := users.UserRegistration(&requestRegister)
	json.NewEncoder(w).Encode(response)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	user := users.GetUser(auth)
	json.NewEncoder(w).Encode(&user)
}

func GetListHotels(w http.ResponseWriter, r *http.Request) {
	hotels := &[]hotels.ResponseHotel{}
	queryset := database.DB.Table("hotels").Select("id, display_name, description, price")

	search_name := r.FormValue("search_name")
	search_description := r.FormValue("search_desc")

	if search_name != "" {
		value := fmt.Sprintf("%%%s%%", r.FormValue("search_name"))
		queryset.Where("display_name ILIKE ?", value)
	}

	if search_description != "" {
		value := fmt.Sprintf("%%%s%%", r.FormValue("search_desc"))
		queryset.Where("description ILIKE ?", value)
	}
	
	queryset.Find(&hotels)
	if queryset.Error != nil {
		json.NewEncoder(w).Encode(utils.HandleResponse("Not found", 404))
		return
	}
	json.NewEncoder(w).Encode(&hotels)
}

func GetInfoHotel(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var hotel hotels.ResponseHotel
	query := database.DB.Table("hotels").Select("id, display_name, description, price").Where("id = ?", params["id"]).First(&hotel, params["id"])
	if query.Error != nil {
		response := utils.HandleResponse("Not found", 404)
		json.NewEncoder(w).Encode(response)
		return
	}
	json.NewEncoder(w).Encode(&hotel)
}


func CreateBooking(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		response := utils.HandleResponse("Unauthorize", 400)
		json.NewEncoder(w).Encode(response)
		return
	}
	isValid, tokenData := utils.ValidateToken(auth)
	if isValid == false {
		response := utils.HandleResponse("Invalid Token", 400)
		json.NewEncoder(w).Encode(response)
		return
	}

	var reqBooking bookings.CreateBookingSerializer
	json.NewDecoder(r.Body).Decode(&reqBooking)

	user_id := tokenData["user_id"].(float64)
	fmt.Println(int(user_id))
	var user users.User
	database.DB.First(&user, int(user_id))

	var account users.ResponseAccount
	database.DB.Table("accounts").Select("id, first_name, last_name, birth_date").Where("user_id = ?", user.ID).Scan(&account)

	responseUser := &users.ResponseUser{
		ID: user.ID,
		Username: user.Username,
		Email: user.Email,
		Account: account,
	}

	var hotel hotels.ResponseHotel
	database.DB.Table("hotels").Select("id, display_name, description, price").Where("id = ?", reqBooking.HotelID).First(&hotel)
	booking := bookings.Booking{
		UserID: user.ID,
		HotelID: hotel.ID}
	
	database.DB.Create(&booking)
	response := &users.ResponseBooking{
		ID: booking.ID,
		User: *responseUser,
		Hotel: hotel,
	}
	json.NewEncoder(w).Encode(&response)
}


func GetListUserBookings(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth == "" {
		response := utils.HandleResponse("Unauthorize", 400)
		json.NewEncoder(w).Encode(response)
		return
	}
	isValid, tokenData := utils.ValidateToken(auth)
	if isValid == false {
		response := utils.HandleResponse("Invalid Token", 400)
		json.NewEncoder(w).Encode(response)
		return
	}

	user_id := tokenData["user_id"].(float64)
	bookings := &[]bookings.ResponseBooking{}

	database.DB.Table("bookings").Select("id, user_Id, hotel_id").Where("user_id = ?", int(user_id)).Scan(&bookings)
	
	json.NewEncoder(w).Encode(&bookings)
}