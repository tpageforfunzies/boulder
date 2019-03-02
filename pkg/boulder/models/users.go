
package models

import (
	"github.com/dgrijalva/jwt-go"
	u "github.com/tpageforfunzies/boulder/pkg/boulder/common"
	"strings"
	"github.com/jinzhu/gorm"
	"os"
	"golang.org/x/crypto/bcrypt"
)

/*
JWT struct
*/
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
	Token string `json:"token";sql:"-"`
}

//Validate incoming user details...
func (user *User) Validate() (may[string] interface{}, bool) {
	if !strings.Contains(account.Email, "@") {
		return u.Message
	}
}