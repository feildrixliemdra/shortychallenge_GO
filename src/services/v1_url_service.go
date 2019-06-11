package services 

import (
	"../objects"
	"../repositories"
	"github.com/jinzhu/copier"
	"crypto/rand"
)
func randStr(strSize int) string {

	const dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
			bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	return string(bytes)
}

type V1UrlService struct {
	urlRepository repositories.V1UrlRepository
	request objects.V1UrlObjectResponse
}

func V1UrlServiceHandler() V1UrlService{
	service := V1UrlService{
		urlRepository: repositories.V1UrlRepositoryHandler(),
	}
	return service
}

func (service *V1UrlService) GetShortenUrl(shorten string) (string, error){
	shortenUrl, err := service.urlRepository.ViewShortenUrl(shorten)
	if nil!= err {
		return "", err
	}


	return shortenUrl, nil
}

func (service *V1UrlService) GetUrlStats(url string) (objects.V1UrlStatsObjectResponse, error){
	shortenUrl,err := service.urlRepository.ViewUrlStats(url)
	if nil!= err {
		return objects.V1UrlStatsObjectResponse{}, err
	}
	result:= objects.V1UrlStatsObjectResponse{}
	copier.Copy(&result,&shortenUrl)
	return result, nil
}


func (service *V1UrlService) CreateNewUrl(InputUrl, ShortenUrl string) (objects.V1NewUrlObjectResponse, error){
	if ShortenUrl==""{
		ShortenUrl:=randStr(6)
		url, err := service.urlRepository.CreateShortenUrl(InputUrl,ShortenUrl)
		if nil != err {
			return objects.V1NewUrlObjectResponse{}, err
		}

		result := objects.V1NewUrlObjectResponse{}
		copier.Copy(&result, &url)

		return result, nil
	}else {
		url, err := service.urlRepository.CreateShortenUrl(InputUrl,ShortenUrl)
		if nil != err {
			return objects.V1NewUrlObjectResponse{}, err
		}

		result := objects.V1NewUrlObjectResponse{}
		copier.Copy(&result, &url)

		return result, nil
	}
}
