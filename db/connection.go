package db

import (
	"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var DSN = "host=localhost user=postgres password=postgres dbname=prueba_go_1 port=5432"
var DB *gorm.DB

func DBConnection(){
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("Conexión exitosa con la base de datos")
	}
}