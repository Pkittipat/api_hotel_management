package api

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()
	log.Fatal(http.ListenAndServe(":8000", router));
}