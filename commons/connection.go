package commons

import (
	"log"

	"github.com/afma0072/aplicacion-inventario-backend/models"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, error := gorm.Open("mysql", "root:@/aplicacion-inventario?charset=utf8&parseTime=true&loc=Local")

	if error != nil {
		log.Fatal(error)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migrando....")

	db.AutoMigrate(&models.Categoria{})
	db.AutoMigrate(&models.Producto{})
	db.AutoMigrate(&models.Transaccion{})
}
