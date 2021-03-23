package api

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"hotel_management/users"
	"encoding/json"
)

func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/api/users/authenticate", AuthenticateUser).Methods("POST")
	router.HandleFunc("/api/users/signup", RegistrationUser).Methods("POST")
	router.HandleFunc("/api/users/me", GetUser).Methods("GET")

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