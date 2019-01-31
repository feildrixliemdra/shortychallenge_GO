package services

import (
	"../objects"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

type V1AuthenticationService struct {
	request objects.V1AuthenticationObjectRequest
}

func V1AuthenticationServiceHandler() (V1AuthenticationService) {
	service := V1AuthenticationService{}
	return service
}

func (service *V1AuthenticationService) Generate(requestObject objects.V1AuthenticationObjectRequest) (objects.V1AuthenticationObjectResponse, error) {

	url := fmt.Sprintf("%s/v2/token", os.Getenv("OAUTH_SERVER_URL"))

	requestBody := map[string]interface{}{
		"grant_type":    "authorization_code",
		"client_secret": os.Getenv("OAUTH_CLIENT_SECRET"),
		"code":          requestObject.Code,
	}

	bytesRepresentation, err := json.Marshal(requestBody)
	if err != nil {
		return objects.V1AuthenticationObjectResponse{}, err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		return objects.V1AuthenticationObjectResponse{}, err
	}

	if resp.StatusCode != 200 {
		return objects.V1AuthenticationObjectResponse{}, errors.New("Invalid authorization!")
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	expiresAt := int(result["expires_at"].(float64))

	resultObject := objects.V1AuthenticationObjectResponse{
		TokenType:    result["token_type"].(string),
		AccessToken:  result["access_token"].(string),
		ExpiresAt:    expiresAt,
		RefreshToken: result["refresh_token"].(string),
	}

	return resultObject, nil

}
