// routes.go
package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Route struct {
	gorm.Model
	Name       string    `json:"name"`
	Grade      string    `json:"grade"`
	UserId     uint      `json:"user_id"` //The user that this contact belongs to
	DateSent   time.Time `json:"date_sent"`
	Type       string    `json:"type"`
	LocationId int       `json:"location_id"`
	Rating     int       `json:"rating"`
	Style      string    `json:"style"`
	ImageUrl   string    `json:"image_url"`
	Comments   []Comment
}
