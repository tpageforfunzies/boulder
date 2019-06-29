// relationship.go
package models

import (
	"github.com/jinzhu/gorm"
)

type Relationship struct {
	gorm.Model
	FollowerID int `json:"follower_id"`
	FollowedID int `json:"followed_id"`
}
