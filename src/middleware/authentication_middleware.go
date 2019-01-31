package middleware

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func (m *DefaultMiddleware) AuthenticationMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {

		authorization := c.GetHeader("Authorization")
		authorizationString := strings.Split(authorization, " ")

		if len(authorizationString) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if strings.ToLower(authorizationString[0]) != "bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := authorizationString[1]

		url := fmt.Sprintf("%s/v2/profile", os.Getenv("OAUTH_SERVER_URL"))

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		req.Header.Add("cache-control", "no-cache")

		res, _ := http.DefaultClient.Do(req)

		if res.StatusCode == 200 {

			defer res.Body.Close()
			body, _ := ioutil.ReadAll(res.Body)

			var result map[string]interface{}
			json.Unmarshal([]byte(string(body)), &result)

			c.Set("user", result)

		} else {

			c.AbortWithStatus(http.StatusUnauthorized)
			return

		}

	}

}
