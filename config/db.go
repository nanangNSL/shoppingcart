package config

import (
	"API-CART/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

  var DB *gorm.DB
func Connect(){
db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=root dbname=gobase port=5432 sslmode=disable"), &gorm.Config{})
if err != nil {
	panic(err)
}
db.AutoMigrate(&models.Cart{
	Model:      gorm.Model{},
	KodeProduk: 0,
	NamaProduk: "",
	Kuantitas:  "",
})
DB = db
}
func InitTestDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=postgres password=root dbname=gobase port=5432 sslmode=disable"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    return db
}
