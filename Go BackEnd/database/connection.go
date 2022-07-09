package database

import (
	"../models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// fungsi koneksi DB
func Connect() {
	connection, err := gorm.Open(mysql.Open("debian-sys-maint:UEokXG7thc0dtQvh@/aino"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
