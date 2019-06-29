// database.go
package services

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	// don't turn this back on until you have gcc in your path
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/joho/godotenv"
	"github.com/tpageforfunzies/boulder/models"
)

var db *gorm.DB
var err error

// Automatically loads from
func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading variables from environment file")
	}
}

func createDatabaseConnection() {
	databaseName := os.Getenv("DATABASE_NAME")
	databaseType := os.Getenv("DATABASE_TYPE")
	databaseHost := os.Getenv("DATABASE_URL")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")

	// postgres is default here, support for sqlite3 below just in case
	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		databaseHost,
		databasePort,
		databaseUser,
		databaseName,
		databasePassword,
	)

	// commented out currently because it hates windows
	// if databaseType == "sqlite3" {
	// 	connectionString = "./database/boulder.db"
	// }

	db, err = gorm.Open(databaseType, connectionString)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	db.Debug().AutoMigrate(&models.User{}, &models.Route{}, &models.Comment{}, &models.Location{}, &models.Relationship{})
}

// Return singleton database connection
func GetDB() *gorm.DB {

	if db != nil {
		return db
	}

	createDatabaseConnection()
	return db
}

// takes the model and gets the string
// name of it, pointer or not
func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

// lowercases and "pluralizes" the model name into a table name
func getTableName(object interface{}) string {
	stringName := fmt.Sprintf("%ss", strings.ToLower(getType(object)))
	return stringName
}

// takes a pointer to a model and an int id and fills the
// model up and returns it
func FindSingleById(object interface{}, id int) interface{} {
	err := GetDB().Table(getTableName(object)).Where("id = ?", id).First(object).Error
	if err != nil {
		return nil
	}
	return object
}
