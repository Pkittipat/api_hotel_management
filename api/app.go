package api

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"hotel_management/users"
	"hotel_management/hotels"
	"hotel_management/database"
	"hotel_management/utils"
	"encoding/json"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users/authenticate", AuthenticateUser).Methods("POST")
	router.HandleFunc("/api/users/signup", RegistrationUser).Methods("POST")
	router.HandleFunc("/api/users/me", GetUser).Methods("GET")
	router.HandleFunc("/api/hotels", GetListHotels).Methods("GET")
	router.HandleFunc("/api/hotels/{id}", GetInfoHotel).Methods("GET")
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