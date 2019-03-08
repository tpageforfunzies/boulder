// locations.go
package models

import (
	"github.com/jinzhu/gorm"
)

type Location struct {
	gorm.Model
	Name string `json:"name"`
	Description string `json:"description"`
	Routes []Route
}