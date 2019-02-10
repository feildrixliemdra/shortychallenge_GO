package database

import "../models"

func Seed() {

	db := GetConnection()
	db.Table("rl_users").Create(&models.User{
		Name:  "Initial User",
		Email: "initial@ralali.com",
		ImageProfile:"/testing-profile.png",
	})

}
