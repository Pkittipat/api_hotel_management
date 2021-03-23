package api

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	DB *gorm.DB
	Router *mux.Router
}

func (app *App) Initialize(user, password, host, port, dbname string) {
	connectionString := 
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok", host, user, password, dbname, port)
	
	var err error
	app.DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	
	if err != nil {
		log.Fatal(err);
	} else {
		fmt.Println("Connection Database...")
	}

	app.Router = mux.NewRouter()
	app.initializeRouters()
}

func (app *App) initializeRouters() {
}

func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.Router));
}