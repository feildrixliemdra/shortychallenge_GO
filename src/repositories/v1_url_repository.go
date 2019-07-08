package repositories

import (
	"time"

	"../database"
	"../models"
	"github.com/jinzhu/gorm"
	//"github.com/jinzhu/copier"
)

type V1UrlRepository struct {
	DB gorm.DB
}

func V1UrlRepositoryHandler() V1UrlRepository {
	repository := V1UrlRepository{DB: *database.GetConnection()}
	return repository
}

func (repository *V1UrlRepository) ViewShortenUrl(url string) (string, error) {
	urlResponse := models.Url{}
	query := repository.DB.Table("rl_urls")
	query = query.Where("shorten_url=?", url)
	query = query.Updates(map[string]interface{}{"redirect_count": gorm.Expr("redirect_count + ?", 1), "last_seen": time.Now()})
	query = query.First(&urlResponse)

	return urlResponse.InputUrl, query.Error
}

func (repository *V1UrlRepository) ViewUrlStats(url string) (models.Url, error) {
	urlResponse := models.Url{}
	query := repository.DB.Table("rl_urls")
	query = query.Where("shorten_url=?", url)
	query = query.First(&urlResponse)

	return urlResponse, query.Error
}

func (repository *V1UrlRepository) CheckShortenUrl(shortenUrl string) error {
	isAvailableUrl := models.Url{}
	query := repository.DB.Table("rl_urls")
	query = query.Where("shorten_url=?", shortenUrl)
	query = query.First(&isAvailableUrl)
	return query.Error
}

func (repository *V1UrlRepository) CreateShortenUrl(inputUrl, shortenUrl string) (models.Url, error) {
	urlModel := models.Url{
		InputUrl:   inputUrl,
		ShortenUrl: shortenUrl,
	}
	query := repository.DB.Create(&urlModel)
	//query=query.NewRecord(urlModel)
	return urlModel, query.Error
}
