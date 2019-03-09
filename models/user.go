// models.go
package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// JWT struct
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//User struct
type User struct {
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token" sql:"-"`
	Routes []Route
	Comments []Comment
}