package database

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rollbar/rollbar-go"
	"go.uber.org/zap"
	"net/url"
	"os"
)

var connection *gorm.DB

func init() {
	godotenv.Load()
}

func Initialize() (*gorm.DB, error) {

	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")

	defaultTimezone := os.Getenv("SERVER_TIMEZONE")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=1&loc=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		url.QueryEscape(defaultTimezone),
	)

	var err error

	connection, err = gorm.Open("mysql", connectionString)
	if nil != err {

		redOutput := color.New(color.FgRed)
		errorOutput := redOutput.Add(color.Bold)

		errorOutput.Println("")
		errorOutput.Println("!!! Warning")
		errorOutput.Println(fmt.Sprintf("Failed connected to database %s", connectionString))
		errorOutput.Println("")

		rollbar.Error(err)

	} else {

		greenOutput := color.New(color.FgGreen)
		successOutput := greenOutput.Add(color.Bold)

		successOutput.Println("")
		successOutput.Println("!!! Info")
		successOutput.Println(fmt.Sprintf("Successfully connected to database %s", connectionString))
		successOutput.Println("")

	}

	zapLog, _ := zap.NewProduction()
	connection.SetLogger(customLogger(zapLog))

	fmt.Println("Connection is created")
	return connection, nil

}

func GetConnection() *gorm.DB {
	if connection == nil {
		fmt.Println("Initialize database")
		connection, _ = Initialize()
	} else {
		fmt.Println("Get connection database")
	}
	return connection
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
