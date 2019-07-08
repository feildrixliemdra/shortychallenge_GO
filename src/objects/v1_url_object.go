package objects

import (
	"time"
)

type V1UrlObjectRequest struct {
	ShortenUrl string `json:"shorten_url"`
}

type V1UrlObjectResponse struct {
	InputUrl string `json:"Location"`
}

type V1UrlStatsObjectResponse struct {
	CreatedAt     time.Time  `json:"start_date"`
	LastSeen      *time.Time `json:"last_seen,omitempty"`
	RedirectCount int        `json:"redirect_count"`
}

type V1UrlStatsObjectRequest struct {
	LastSeen      time.Time `json:"last_seen"`
	RedirectCount int       `json:"redirect_count"`
}

type V1NewUrlObjectRequest struct {
	InputUrl   string `json:"input_url"`
	ShortenUrl string `json:"shorten_url"`
}

type V1NewUrlObjectResponse struct {
	ShortenUrl string `json:"shorten_url"`
}
