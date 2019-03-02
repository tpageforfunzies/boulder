// database.go
package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/joho/godotenv"
	"log"
	"os"
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

	if databaseType == "sqlite3" {
		connectionString = "./database/boulder.db"
	}



	db, err = gorm.Open(databaseType, connectionString)
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	db.Debug().AutoMigrate(&User{})
}

// Return singleton database connection
func GetDB() *gorm.DB {

	if db != nil {
		return db
	}

	createDatabaseConnection()
	return db
}
