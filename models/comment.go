// comments.go
package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserId uint `json:"user_id"`
	RouteId uint `json:"route_id"`
	Content string `json:"content"`
}