package main

import (
	"./controllers"
	"./database"
	"./middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rollbar/rollbar-go"
	"log"
	"os"
)

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	if os.Getenv("APP_ENV") == "production" {
		rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
		rollbar.SetEnvironment(os.Getenv("APP_ENV"))
		rollbar.WrapAndWait(startApp)
	} else {
		startApp()
	}

}

func startApp() {

	defaultMiddleware := middleware.DefaultMiddleware{}

	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())

	controllers.V1UserControllerHandler(router)
	controllers.V1AuthenticationControllerHandler(router)
	controllers.V2UserControllerHandler(router)

	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println(serverString)

	// run database migration
	database.Migrate()

	router.Run(serverString)

}
