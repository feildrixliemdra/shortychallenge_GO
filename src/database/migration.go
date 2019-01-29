package database

import (
	"../models"
)

func Migrate() {

	db := GetConnection()
	db.AutoMigrate(&models.User{})

}
