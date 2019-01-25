package repositories

import (
	"../database"
	"../models"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
)

type V2UserRepository struct {
	DB gorm.DB
}

func V2UserRepositoryHandler() (V2UserRepository) {
	repository := V2UserRepository{DB: *database.GetConnection()}
	return repository
}

func (repository *V2UserRepository) UpdateById(id int, userData interface{}) (models.User, error) {

	userModel := models.User{}
	copier.Copy(&userModel, &userData)

	query := repository.DB.Table("rl_users")
	query = query.Where("id=?", id)
	query = query.Updates(userModel)
	query = query.Scan(&userModel)

	return userModel, query.Error

}
