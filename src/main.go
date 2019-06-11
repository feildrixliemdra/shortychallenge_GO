package main

import (
	"./controllers"
	"./database"
	"./docs"
	"./middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/rollbar/rollbar-go"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"os"
)

func init() {
	if godotenv.Load() != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	cmdString := command()

	if cmdString == "" {
		cmdString = "serve"
	}

	if cmdString == "serve" && os.Getenv("APP_ENV") == "production" {
		startAppWithRollbar()
	} else if cmdString == "serve" && os.Getenv("APP_ENV") != "production" {
		startApp()
	} else if cmdString == "seed" {
		database.Seed()
	} else if cmdString == "migrate" {
		database.Migrate()
	}

}

func startAppWithRollbar() {
	rollbar.SetToken(os.Getenv("ROLLBAR_TOKEN"))
	rollbar.SetEnvironment(os.Getenv("APP_ENV"))
	rollbar.WrapAndWait(startApp)
}

func startApp() {

	defaultMiddleware := middleware.DefaultMiddleware{}

	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())

	controllers.V1UserControllerHandler(router)
	controllers.V1AuthenticationControllerHandler(router)
	controllers.V2UserControllerHandler(router)
	controllers.V1UrlControllerHandler(router)
	// start documentations here
	docs.SwaggerInfo.Title = "Golang Boilerplate"
	docs.SwaggerInfo.Description = "Golang boilerplate endpoint documentations"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("SERVER_ADDRESS"), os.Getenv("SERVER_PORT"))
	docs.SwaggerInfo.BasePath = "/"

	router.GET("/documentations/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println(serverString)

	router.Run(serverString)

}

func command() string {
	args := os.Args[1:]

	if len(args) > 0 {
		return args[0]
	}
	return ""
}
