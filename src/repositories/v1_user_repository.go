package repositories

import (
	"../database"
	"../models"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type V1UserRepository struct {
	DB gorm.DB
}

func V1UserRepositoryHandler() (V1UserRepository) {
	repository := V1UserRepository{DB: *database.GetConnection()}
	return repository
}

func (repository *V1UserRepository) GetById(id int) (models.User, error) {

	userResponse := models.User{}

	query := repository.DB.Table("rl_users")
	query = query.Where("id=?", id)
	query = query.First(&userResponse)

	return userResponse, query.Error

}

func (repository *V1UserRepository) UpdateById(id int, userData interface{}) (models.User, error) {

	userModel := models.User{}
	copier.Copy(&userModel, &userData)

	query := repository.DB.Table("rl_users")
	query = query.Where("id=?", id)
	query = query.Updates(userModel)
	query = query.Scan(&userModel)

	return userModel, query.Error

}
