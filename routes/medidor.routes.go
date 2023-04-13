package routes

import (
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"net/http"

	"time"

	"github.com/adriangarcia1984/go-postgress/database"
	"github.com/adriangarcia1984/go-postgress/models"
	"github.com/google/uuid"
	"github.com/gorilla/mux"

	_ "github.com/adriangarcia1984/go-postgress/docs"
)


// ShowMedidor godoc
// @Summary      Show an medidor
// @Description  get medidor por id
// @Tags         medidores
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "medidor ID"
// @Success      200  {object}  models.Medidor
// @Failure      400  
// @Failure      404  
// @Failure      500  
// @Router       /medidores/{id} [get]
func GetMedidorHandler(w http.ResponseWriter, r *http.Request) {

	var medidor models.Medidor
	params:=mux.Vars(r)
	//fmt.Println(params["id"])

	database.DB.First(&medidor, params["id"])

	if medidor.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)//400
		w.Write([]byte("medidor no encontrado"))
		return
	}

	json.NewEncoder(w).Encode(&medidor)

	w.Write([]byte("getmedidorhandler"))
}

//ShowAllMedidores godoc
// @Summary      Show all medidores
// @Description  get medidores
// @Tags         medidores
// @Accept       json
// @Produce      json
// @Success      200  {array}  models.Medidores
// @Failure      400  
// @Failure      404  
// @Failure      500  
// @Router       /medidores [get]
func GetMedidoresHandler(w http.ResponseWriter, r *http.Request) {
	
	var medirores []models.Medidor

	database.DB.Find(&medirores)

	json.NewEncoder(w).Encode(&medirores)

	w.Write([]byte("medidores consultados"))

}

//create an medidor godoc
// @Summary      crea un medidor en la base de datos
// @Description  crea un medidor en la base de datos
// @Tags         medidores
// @Accept       json
// @Produce      json
// @Param        medidor   body      models.Medidor  true  "crea un nuevo medidor"
// @Success      200  {object}  models.Medidor
// @Failure      400  
// @Failure      404  
// @Failure      500  
// @Router       /medidores [post]
func PostMedidorHandler(w http.ResponseWriter, r *http.Request) {
	
	var medidor, newMedidor models.Medidor

	medidor = models.Medidor{
		CreatedAt: 	time.Now(),
		Id: 		uuid.New().ClockSequence(),
	}

	json.NewDecoder(r.Body).Decode(&medidor)
	
	//fmt.Println("body ", &medidor)

	database.DB.Where(&models.Medidor{Brand: medidor.Brand, Serial: medidor.Serial}).Find(&newMedidor)

	//fmt.Println("consulta ", &newMedidor)

	if newMedidor.Id !=0 {
		w.WriteHeader(http.StatusMethodNotAllowed)//405
		w.Write([]byte("no se puede crear el medidor, ya hay uno registrado con los datos ingresados"))
		return
	}
	createdMedidor:=database.DB.Save(&medidor)

	err:= createdMedidor.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)//400
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&medidor)
	w.Write([]byte("medidor creado con exito"))
}

//update an medidor godoc
// @Summary      actualiza un medidor en la base de datos
// @Description  actualiza un medidor en la base de datos
// @Tags         medidores
// @Router       /medidores/{id} [put]
// @Accept       json
// @Produce      json
// @Param        id   path  int  true  "medidor ID"
// @Param        medidor   body      models.Medidor  false  "actualiza un nuevo medidor, los campos como brand, serial, id, createdAt, no son editables"
// @Success      200  {object}  models.Medidor
// @Failure      400  
// @Failure      404  
// @Failure      500  
func UpdateMedidorHandler(w http.ResponseWriter, r *http.Request) {

	//var  prueba models.Medidor
	var medidor, newMedidor models.Medidor
	
	json.NewDecoder(r.Body).Decode(&newMedidor)
	
	//fmt.Println("new medidor: ",newMedidor)

	params:=mux.Vars(r)
	
	database.DB.First(&medidor, params["id"])

	fmt.Println("medidor: ",medidor)
	
	if medidor.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)//400
		w.Write([]byte("medidor not found"))
		return
	}

	if (newMedidor.Brand !="" || newMedidor.Serial!="" )  {
		w.WriteHeader(http.StatusBadRequest)//400
		w.Write([]byte("el dato que intenta modificar esta bloqueado"))
		return
	}

	if (newMedidor.Address !="" && medidor.IsActive!=false )  {
		w.WriteHeader(http.StatusBadRequest)//400
		w.Write([]byte("el medidor que intenta modificar se encuentra activo en una direccion, por favor, desactive el medidor"))
		return
	}

	//fmt.Println(newMedidor.IsActive)

	database.DB.Model(&medidor).Updates(models.Medidor{

		IsActive: 			newMedidor.IsActive,
		Address: 			newMedidor.Address,
		Lines: 				newMedidor.Lines,
		InstallationDate:	newMedidor.InstallationDate,
		RetirementDate: 	newMedidor.RetirementDate,
	
	})


	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(&medidor)
	w.Write([]byte("medidor actualizado con exito"))
	//json.NewEncoder(w).Encode(&newMedidor)
}

//delete an medidor godoc
// @Summary      elimina un medidor en la base de datos por su id
// @Description  elimina un medidor en la base de datos por su id
// @Tags         medidores
// @Produce      json
// @Param        id   path      int  true  "medidor ID"
// @Success      200  string  "medidor eliminado"
// @Failure      400  
// @Failure      404  
// @Failure      500  
// @Router       /medidores/{id} [delete]
func DeleteMedidorHandler(w http.ResponseWriter, r *http.Request) {

	var medidor models.Medidor
	params:=mux.Vars(r)

	database.DB.First(&medidor, params["id"])

	if medidor.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("medidor not found"))
		return
	}

	if medidor.IsActive==true{
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("no se puede eliminar el medidor, se encuentra activo"))
		return
	}
	//solamente cambia el campo deleted_at con una fecha y no lo lista pero no lo elimina
	//database.DB.Delete(&medidor)

	//con este metodo lo elimina por completo de la tabla
	database.DB.Unscoped().Delete(&medidor)


	w.Write([]byte("DeleteMedidorHandler"))
}