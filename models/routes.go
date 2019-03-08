// routes.go
package models

import (
	"github.com/jinzhu/gorm"
)

type Route struct {
	gorm.Model
	Name string `json:"name"`
	Grade string `json:"grade"`
	UserId uint `json:"user_id"` //The user that this contact belongs to
}