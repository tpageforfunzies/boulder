// models.go
package models

import (
	"github.com/dgrijalva/jwt-go"
	u "github.com/tpageforfunzies/boulder/pkg/boulder/common"
	"strings"
	"github.com/jinzhu/gorm"
	"os"
	"golang.org/x/crypto/bcrypt"
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
	Token string `json:"token";sql:"-"`
}

//Validate incoming user details
// Returns message, ok/success
func (user *User) Validate() (map[string] interface{}, bool) {

	// basic format check
	if !strings.Contains(user.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(user.Password) < 6 {
		return u.Message(false, "Cmon my doggie"), false
	}

	// Set up an object just in case
	check := &User{}

	// see if they're already in there and if so, put in check
	err := GetDB().Table("users").Where("email = ?", user.Email).First(check).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, string(err.Error())), false
	}

	if check.Email != "" {
		return u.Message(false, "Someone is already using that email"), false
	}

	// not already in db
	return u.Message(false, "i believe you"), true
}

func (user *User) Create() (map[string] interface{}) {

	// make sure they're not in db already
	dupe, ok := user.Validate()
	if !ok {
		return dupe
	}

	// hash the hell out of the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	// ugh data mapper
	GetDB().Create(user)

	// should have an id now
	if user.ID <= 0 {
		return u.Message(false, "don't have an ID")
	}

	//make the token
	tkn := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tkn)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	// get it up out of here
	user.Password = ""

	resp := u.Message(true, "User created!")
	resp["user"] = user
	return resp
}

func Login(email, password string) (map[string]interface{}) {

	user := &User{}
	err := GetDB().Table("users").Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Email address not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match
		return u.Message(false, "Invalid login credentials. Please try again")
	}
	//Worked Logged In
	user.Password = ""

	//Create JWT token
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["user"] = user
	return resp
}

func GetUser(u uint) *User {

	usr := &User{}
	GetDB().Table("users").Where("id = ?", u).First(usr)
	if usr.Email == "" { //User not found
		return nil
	}

	usr.Password = ""
	return usr
}