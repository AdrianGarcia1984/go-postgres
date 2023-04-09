package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/adriangarcia1984/go-postgress/database"
	"github.com/adriangarcia1984/go-postgress/models"
	"github.com/gorilla/mux"
)

func ActiveMedidor(w http.ResponseWriter, r *http.Request) {

	var medidores models.Medidores

	params := mux.Vars(r)

	isActive := false
	if params["is_active"] == "true" {
		isActive = true
	}

	database.DB.Where("is_active = ?", isActive).Find(&medidores)

	json.NewEncoder(w).Encode(&medidores)
	w.Write([]byte("active medidor"))
}

func BrandMedidor(w http.ResponseWriter, r *http.Request) {
	var medidor models.Medidor
	params := mux.Vars(r)

	Serial:=params["serial"]
	brand:= strings.Replace(params["brand"], "_", " ", -1)  

	fmt.Println("brand: ", brand, " serial: ", Serial)

	database.DB.Where("serial = ?", Serial).Find(&medidor)

	fmt.Println(medidor.InstallationDate)
	w.Write([]byte("fecha de instalacion: "))
	json.NewEncoder(w).Encode(&medidor.InstallationDate)
}
