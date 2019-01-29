package objects

type V1UserObjectResponse struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	ImageProfile string `json:"image_profile"`
}

type V1UserObjectRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
