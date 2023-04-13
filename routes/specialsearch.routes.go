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


//medidor activo godoc
// @Summary      busca un medidor en la base de datos con su id y develve si se encuentra activo
// @Description  busca un medidor en la base de datos con su id y devielve si se encuentra activo
// @Tags         medidores/active
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "medidor ID"

// @Success      200  string  "medidor eliminado"
// @Failure      400  
// @Failure      404  
// @Failure      500  
// @Router       /medidores/active/{is_Active} [get]
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


//medidor por marca y serial godoc
// @Summary      busca un medidor en la base de datos con su marca y serial y develve si se la fecha de instalacion
// @Description  busca un medidor en la base de datos con su marca y serial y develve si se la fecha de instalacion
// @Tags         medidores/active
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "medidor ID"

// @Success      200  string  "medidor eliminado"
// @Failure      400  
// @Failure      404  
// @Failure      500  
// @Router       /medidores/search/{brand}/{serial} [get]
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
