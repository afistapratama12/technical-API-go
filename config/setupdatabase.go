package config

import (
	"golang-api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@(localhost)/golang_api?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("gagal koneksi database")
	}

	db.AutoMigrate(&models.Item{})
	db.AutoMigrate(&models.Order{})

	return db
}
