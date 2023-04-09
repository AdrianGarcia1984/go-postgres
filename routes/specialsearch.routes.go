package routes

import (
	"encoding/json"
	//"fmt"
	"net/http"

	"github.com/adriangarcia1984/go-postgress/database"
	"github.com/adriangarcia1984/go-postgress/models"
	"github.com/gorilla/mux"
)

func ActiveMedidor(w http.ResponseWriter, r *http.Request) {

	var medidores models.Medidores

	params:=mux.Vars(r)

	isActive:=false
	if params["is_active"]=="true"{
		isActive=true
	}

	database.DB.Where("is_active = ?", isActive).Find(&medidores)

	json.NewEncoder(w).Encode(&medidores)
	w.Write([]byte("active medidor"))
}