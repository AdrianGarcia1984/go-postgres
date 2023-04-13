package main

import (
	"net/http"

	"github.com/adriangarcia1984/go-postgress/database"
	"github.com/adriangarcia1984/go-postgress/models"
	"github.com/adriangarcia1984/go-postgress/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"

	_ "github.com/adriangarcia1984/go-postgress/docs"

	//importando archivos de swagger
	httpSwagger "github.com/swaggo/http-swagger" // http-swagger middleware
	


)

// @title API Medidores de Energia
// @version 1.0.1
// @description Esta api se encarga de llevar el control de los medidores de energia de la empresa enerBit, donde se presenta el reto de registrar y hacer seguimiento de los medidores de energía que se han instalado en los predios de nuestros clientes,  para efectos de pruebas en modo local se manejara la url localhost:3000, recuerde primero inicialiar la base de datos, tal como se solicito en la prueba tecnica, se realizo un container en docker el cual tiene la base de datos, los archivos de ejemplo para la conexion a la base de datos se especifican en el archivo env.example, estos deben guardarse como variables de entorno. para mas info consultar archivo readme.md
//

// @Host      localhost:3000
// @BasePath  /

// @Summary      migracion de tablas en postgres
// @Description  migracion de tablas en postgres
func main() {

	 database.DBConection()
	 database.DB.AutoMigrate(models.Medidor{})

	r:=mux.NewRouter()

	// endpoints
	r.HandleFunc("/", routes.HomeHandler)
	r.HandleFunc("/medidores", routes.GetMedidoresHandler).Methods("GET")
	r.HandleFunc("/medidores/{id}", routes.GetMedidorHandler).Methods("GET")
	r.HandleFunc("/medidores", routes.PostMedidorHandler).Methods("POST")
	r.HandleFunc("/medidores/{id}", routes.DeleteMedidorHandler).Methods("DELETE")
	r.HandleFunc("/medidores/{id}", routes.UpdateMedidorHandler).Methods("PUT")
	r.HandleFunc("/medidores/active/{is_active}", routes.ActiveMedidor).Methods("GET")
	r.HandleFunc("/medidores/search/{brand}/{serial}", routes.BrandMedidor).Methods("GET")
	
	//endpoint para swagger
	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)


	handler := cors.Default().Handler(r)

	http.ListenAndServe(":3000", handler)

}
