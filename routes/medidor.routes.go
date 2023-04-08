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
)

//listar un medidor
func GetMedidorHandler(w http.ResponseWriter, r *http.Request) {

	var medidor models.Medidor
	params:=mux.Vars(r)
	fmt.Println(params["id"])

	database.DB.First(&medidor, params["id"])

	if medidor.Id == 0 {
		w.Write([]byte("user not found"))
		return
	}

	json.NewEncoder(w).Encode(&medidor)

	w.Write([]byte("getmedidorhandler"))
}

//listar todos los medidores
func GetMedidoresHandler(w http.ResponseWriter, r *http.Request) {
	
	var medirores []models.Medidor

	database.DB.Find(&medirores)

	json.NewEncoder(w).Encode(&medirores)

	w.Write([]byte("medidores consultados"))

}

//post function
func PostMedidorHandler(w http.ResponseWriter, r *http.Request) {
	
	var medidor, newMedidor models.Medidor

	medidor = models.Medidor{
		CreatedAt: 	time.Now(),
		Id: 		uuid.New().ClockSequence(),
	}

	json.NewDecoder(r.Body).Decode(&medidor)
	
	fmt.Println("body ", &medidor)

	database.DB.Where(&models.Medidor{Brand: medidor.Brand, Serial: medidor.Serial}).Find(&newMedidor)

	fmt.Println("consulta ", &newMedidor)

	if newMedidor.Id !=0 {
		w.WriteHeader(http.StatusMethodNotAllowed)
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

//update function
func UpdateMedidorHandler(w http.ResponseWriter, r *http.Request) {

	//var  prueba models.Medidor
	var medidor, newMedidor models.Medidor
	
	json.NewDecoder(r.Body).Decode(&newMedidor)
	
	// reqbody,_:= ioutil.ReadAll(r.Body)	
	// json.Unmarshal(reqbody,&prueba)
	
	// fmt.Println("prueba: ",reqbody)
	
	fmt.Println("new medidor: ",newMedidor)

	params:=mux.Vars(r)
	
	database.DB.First(&medidor, params["id"])

	fmt.Println("medidor: ",medidor)
	
	if medidor.Id == 0 {
		w.Write([]byte("medidor not found"))
		return
	}

	if (newMedidor.Brand !="" || newMedidor.Serial!="" )  {
		w.WriteHeader(http.StatusBadRequest)//400
		w.Write([]byte("el dato que intenta modificar esta bloqueado"))
		return
	}

	if (newMedidor.Address !="" && medidor.IsActive!=false )  {
		w.Write([]byte("el medidor que intenta modificar se encuentra activo en una direccion, por favor, desactive el medidor"))
		w.WriteHeader(http.StatusBadRequest)//400
		return
	}

	fmt.Println(newMedidor.IsActive)

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
	json.NewEncoder(w).Encode(&newMedidor)
}

//delete medidor
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