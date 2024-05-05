package routes

import (
	"encoding/json"
	"net/http"

	"github.com/N13AA/go_apirest/db"
	"github.com/N13AA/go_apirest/models"
	"github.com/gorilla/mux"
)

func Getubication(w http.ResponseWriter, r *http.Request) {
	type Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	// Crear una instancia de Location con los valores de latitud y longitud
	ubi := Location{
		Latitude:  13.678750,
		Longitude: -89.266849,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ubi)

}
func Postubication(w http.ResponseWriter, r *http.Request) {

	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	db.DB.Create(&user)
	json.NewEncoder(w).Encode(&user)

}
func Getuser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(&user)
}
