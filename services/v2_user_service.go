package services

import (
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"
	"ralali.com/objects"
	"ralali.com/repositories"
)

type V2UserService struct {
	request        objects.V2UserObjectResponse
	userRepository repositories.V2UserRepository
}

func V2UserServiceHandler(db *gorm.DB) (V2UserService) {
	service := V2UserService{
		userRepository: repositories.V2UserRepositoryHandler(db),
	}
	return service
}

func (service *V2UserService) UpdateById(id int, requestObject objects.V2UserObjectRequest) (objects.V2UserObjectResponse, error) {

	user, err := service.userRepository.UpdateById(id, requestObject)
	if nil != err {
		return objects.V2UserObjectResponse{}, err
	}

	result := objects.V2UserObjectResponse{}
	copier.Copy(&result, &user)

	return result, nil

}