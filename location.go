package main

import (
	"net/http"

	"github.com/N13AA/go_apirest/db"
	"github.com/N13AA/go_apirest/models"
	"github.com/N13AA/go_apirest/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBconnection()
	db.DB.AutoMigrate(models.User{})
	url := mux.NewRouter()
	url.HandleFunc("/", routes.Homehandler)
	url.HandleFunc("/location", routes.Getubication).Methods("GET")
	url.HandleFunc("/location/{id}", routes.Getuser).Methods("GET")
	url.HandleFunc("/location", routes.Postubication).Methods("POST")
	http.ListenAndServe(":8080", url)
}
