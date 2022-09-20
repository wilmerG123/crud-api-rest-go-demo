package commons

import (
	"crud-api-rest-go-demo/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:00000000@tcp(localhost:3306)/goweb_db?charset=utf8")
	if error != nil {
		log.Fatal(error)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()
	log.Println("Migrando....")
	db.AutoMigrate(&models.Persona{})
}
