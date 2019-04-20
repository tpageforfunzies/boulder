// user_service.go
package services

import (
    "github.com/tpageforfunzies/boulder/models"
    "github.com/dgrijalva/jwt-go"
	"strings"
	"github.com/jinzhu/gorm"
	"os"
	"golang.org/x/crypto/bcrypt"
)

//Validate incoming user details
// Returns message, ok/success
func ValidateUser(user *models.User) (string, bool) {

	// basic format check
	if !strings.Contains(user.Email, "@") {
		return "Email address is required", false
	}

	if len(user.Password) < 6 {
		return "Cmon my doggie", false
	}

	// Set up an object just in case
	check := &models.User{}

	// see if they're already in there and if so, put in check
	err := GetDB().Table("users").Where("email = ?", user.Email).First(check).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return "Error checking for email", false
	}

	if check.Email != "" {
		return "Someone is already using that email", false
	}

	// not already in db
	return "i believe you", true
}

// func UpdateUser(user *models.User) (map[string] interface{}) {
	
// }

func CreateUser(user *models.User) (string, *models.User) {

	// make sure they're not in db already
	validationResult, ok := ValidateUser(user)
	if !ok {
		return validationResult, user
	}

	// hash the hell out of the password
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)

	GetDB().Create(user)

	// should have an id now
	if user.ID == 0 {
		return "don't have an ID", user
	}

	//make the token
	tkn := &models.Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tkn)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString

	// get it up out of here
	user.Password = ""

	return "User created!", user
}

func Login(email, password string) (string, *models.User) {

	user := &models.User{}
	err := GetDB().Table("users").Where("email = ?", email).Preload("Comments").Preload("Routes").First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "Email address not found", user
		}
		return "Connection error. Please retry", user
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match
		return "Invalid login credentials. Please try again", user
	}
	//Worked Logged In
	user.Password = ""

	//Create JWT token
	tk := &models.Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	user.Token = tokenString //Store the token in the response

	return "Logged In", user
}

func GetUserById(id int) *models.User {
	user := &models.User{}
	err := GetDB().Table("users").Preload("Comments").Preload("Routes").Find(user, id).Error
	if err != nil {
		return nil
	}

	if user.Email == "" { //User not found
		return nil
	}

	// Gorm will get the token out
	user.Password = ""
	return user
}

func GetAllUsers() []*models.User {
	users := make([]*models.User, 0)
	err := GetDB().Table("users").Preload("Comments").Preload("Routes").Find(&users).Error
	if err != nil {
		return nil
	}
	return users
}