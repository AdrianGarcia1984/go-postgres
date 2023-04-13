package database

import (

	"log"
	"os"

	"github.com/adriangarcia1984/go-postgress/env"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var host = env.EnvConfig("HOST",os.Getenv("HOST"))
var user = os.Getenv("USER")
var password = os.Getenv("PASSWORD")
var dbname = os.Getenv("DBNAME")
var port = os.Getenv("PORT")

var DSNENV = ("host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port)


var DB *gorm.DB

func DBConection() {

	var err error
	DB, err = gorm.Open(postgres.Open(DSNENV), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("db connected")
}
