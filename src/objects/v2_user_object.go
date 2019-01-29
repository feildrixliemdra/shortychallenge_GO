package objects

import "time"

type V2UserObjectResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type V2UserObjectRequest struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
