package config

import (
	"golang-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	// setup connecting to database mysql "golang_api" with user root and no password
	db, err := gorm.Open("mysql", "root:@(localhost)/golang_api?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("gagal koneksi database")
	}

	//auto migrate to databases
	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Order{})

	return db
}
