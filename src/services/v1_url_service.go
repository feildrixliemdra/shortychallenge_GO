package services

import (
	"crypto/rand"
	"regexp"

	"../objects"
	"../repositories"
	"github.com/jinzhu/copier"
)

func isAlphaNum(str string) bool {
	re := regexp.MustCompile("^([A-Za-z]+[0-9]|[0-9]+[A-Za-z])[A-Za-z0-9]+$")
	if re.MatchString(str) {
		return true
	} else {
		return false
	}
}

func randStr() string {
	const strSize = 6
	const dictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	var bytes = make([]byte, strSize)
	rand.Read(bytes)
	for k, v := range bytes {
		bytes[k] = dictionary[v%byte(len(dictionary))]
	}
	str := string(bytes)

	if isAlphaNum(str) {
		return str
	}
	return randStr()
}

type V1UrlService struct {
	urlRepository repositories.V1UrlRepository
	request       objects.V1UrlObjectResponse
}

func V1UrlServiceHandler() V1UrlService {
	service := V1UrlService{
		urlRepository: repositories.V1UrlRepositoryHandler(),
	}
	return service
}

func (service *V1UrlService) GetShortenUrl(shorten string) (string, error) {
	shortenUrl, err := service.urlRepository.ViewShortenUrl(shorten)
	if nil != err {
		return "", err
	}

	return shortenUrl, nil
}

func (service *V1UrlService) GetUrlStats(url string) (objects.V1UrlStatsObjectResponse, error) {
	shortenUrl, err := service.urlRepository.ViewUrlStats(url)
	if nil != err {
		return objects.V1UrlStatsObjectResponse{}, err
	}
	result := objects.V1UrlStatsObjectResponse{}

	if shortenUrl.RedirectCount == 0 {
		copier.Copy(&result, &shortenUrl)
		result = objects.V1UrlStatsObjectResponse{result.CreatedAt, nil, result.RedirectCount}
		return result, nil
	}
	copier.Copy(&result, &shortenUrl)
	return result, nil
}

func (service *V1UrlService) CheckShortenUrl(ShortenUrl string) (bool, error) {
	err := service.urlRepository.CheckShortenUrl(ShortenUrl)

	if nil != err {

		return false, err

	}

	return true, err

}

func (service *V1UrlService) CreateNewUrl(InputUrl, ShortenUrl string) (objects.V1NewUrlObjectResponse, error) {
	if ShortenUrl == "" {
		ShortenUrl := randStr()
		url, err := service.urlRepository.CreateShortenUrl(InputUrl, ShortenUrl)
		if nil != err {
			return objects.V1NewUrlObjectResponse{}, err
		}

		result := objects.V1NewUrlObjectResponse{}
		copier.Copy(&result, &url)

		return result, nil
	} else {
		url, err := service.urlRepository.CreateShortenUrl(InputUrl, ShortenUrl)
		if nil != err {
			return objects.V1NewUrlObjectResponse{}, err
		}

		result := objects.V1NewUrlObjectResponse{}
		copier.Copy(&result, &url)

		return result, nil
	}
}
