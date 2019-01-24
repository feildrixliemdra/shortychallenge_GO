package main

import (
	"./controllers"
	"./middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rollbar/rollbar-go"
	"go.uber.org/zap"
	"github.com/fatih/color"
	"log"
	"net/url"
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

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")

	defaultTimezone := os.Getenv("SERVER_TIMEZONE")

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1&loc=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		url.QueryEscape(defaultTimezone),
	)

	db, err := gorm.Open("mysql", connection)
	if nil != err {
		
		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)
		
		errorOutput.Println("")
		errorOutput.Println("!!! Warning")
		errorOutput.Println(fmt.Sprintf("Failed connected to database %s", connection))
		errorOutput.Println("")

		rollbar.Error(err)

	} else {
		

		greenOutput := color.New(color.FgGreen)
		successOutput := greenOutput.Add(color.Bold)
		
		successOutput.Println("")
		successOutput.Println("!!! Info")
		successOutput.Println(fmt.Sprintf("Successfully connected to database %s", connection))
		successOutput.Println("")

	}

	zapLog, _ := zap.NewProduction()
	db.SetLogger(customLogger(zapLog))

	defer db.Close()

	defaultMiddleware := middleware.DefaultMiddleware{
		DB: db,
	}

	router := gin.Default()
	router.Use(defaultMiddleware.CORSMiddleware())

	controllers.V1UserControllerHandler(router, db)
	controllers.V2UserControllerHandler(router, db)

	serverHost := os.Getenv("SERVER_ADDRESS")
	serverPort := os.Getenv("SERVER_PORT")

	serverString := fmt.Sprintf("%s:%s", serverHost, serverPort)
	fmt.Println(serverString)

	router.Run(serverString)

}

func customLogger(zap *zap.Logger) *customLoggerStruct {
	return &customLoggerStruct{
		zap: zap,
	}
}

type customLoggerStruct struct {
	zap *zap.Logger
}

func (l *customLoggerStruct) Print(values ...interface{}) {
	var additionalString = ""
	for _, item := range values {
		if _, ok := item.(string); ok {
			additionalString = additionalString + fmt.Sprintf("\n%v", item)
		}
		if err, ok := item.(*mysql.MySQLError); ok {
			err.Message = err.Message + additionalString
			if os.Getenv("APP_ENV") == "production" {
				rollbar.Error(err)
			}
		}
	}
}
