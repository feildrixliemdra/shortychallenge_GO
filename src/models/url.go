package models

import(
	"time"
)

type Url struct {
	ID uint `gorm:"primary_key"`
	InputUrl string `gorm:"column:input_url;required" json:"input_url"`
	ShortenUrl string `gorm:"column:shorten_url;unique" json:"shorten_url"`
	RedirectCount int `gorm:"column:redirect_count" sql:"default:0" json:"redirect_count"`
	CreatedAt *time.Time `sql:"index" json:"created_at"`
	DeletedAt *time.Time `sql:"index"`
	LastSeen *time.Time	`json:"last_seen"`
}

func (Url) TableName() string {
	return "rl_urls"
}