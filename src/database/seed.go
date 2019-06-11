package database

import "../models"

func Seed() {

	db := GetConnection()
	// db.Table("rl_users").Create(&models.User{
	// 	Name:  "Initial User",
	// 	Email: "initial@ralali.com",
	// 	ImageProfile:"/testing-profile.png",
	// })

	db.Table("rl_urls").Create(&models.Url{
		InputUrl:  "http://google.com",
		ShortenUrl: "google",
	})

}
