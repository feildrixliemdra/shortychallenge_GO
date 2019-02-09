package repositories

import (
	"../database"
	"../models"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	if godotenv.Load("../.env") != nil {
		log.Fatal("Error loading .env file")
	}
	m.Run()
}

func TestV1UserRepository_UpdateById(t *testing.T) {

	db := database.GetConnection()
	tx := db.Begin()

	// seeds data first
	tx.Table("rl_users").Create(&models.User{
		ID:           1,
		Name:         "Testing name",
		ImageProfile: "/testing-image.jpg",
		Email:        "mq.aashari@gmail.com",
	})

	newData := models.User{
		Name: "Andi Testing",
	}

	repository := V1UserRepositoryHandler()
	repository.DB = *tx

	_, err := repository.UpdateById(1, newData)

	if err != nil {
		t.Error("Failed get data")
	}

	tx.Rollback()

}

func TestV1UserRepository_GetById(t *testing.T) {

	db := database.GetConnection()
	tx := db.Begin()

	// seeds data first
	tx.Table("rl_users").Create(&models.User{
		ID:           1,
		Name:         "Testing name",
		ImageProfile: "/testing-image.jpg",
		Email:        "mq.aashari@gmail.com",
	})

	repository := V1UserRepositoryHandler()
	repository.DB = *tx

	_, err := repository.GetById(1)

	if err != nil {
		t.Error(fmt.Sprintf("Failed get data: %s", err))
	}

	tx.Rollback()

}
