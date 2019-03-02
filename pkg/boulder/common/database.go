// database.go
package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
  	_ "github.com/jinzhu/gorm/dialects/postgres"
  	"sync"
)

var db *gorm.DB

// Automatically loads from 
func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading variables from environment file")
	}
}

func createDatabaseConnection() {
	databaseName := os.Getenv("DATABASE")
	databaseType := os.Getenv("DATABASE_TYPE")
	databaseHost := os.Getenv("DATABASE_URL")
	databasePort := os.Getenv("DATABASE_PORT")
	databaseUser := os.Getenv("DATABASE_USER")
	databasePassword := os.Getenv("DATABASE_PASSWORD")

	db = gorm.Open(databaseType, 
		fmt.Sprintf(
			"host=%s port=%s user=%s dbname=%s password=%s",
			databaseHost,
			databasePort,
			databaseUser,
			databaseName,
			databasePassword
		)
	)

	return db
}

// Return singleton database connection
func GetDBConnection() *gorm.DB {

	if db != nil {
		return db
	}

	return createDatabaseConnection()
}