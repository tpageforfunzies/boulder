// auth.go
package handlers

import (
	u "github.com/tpageforfunzies/boulder/common"
	"github.com/tpageforfunzies/boulder/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
)


func CreateUser(c *gin.Context) {

	user := &models.User{}
	// throw this bitch up in that object
	err := json.NewDecoder(c.Request.Body).Decode(user)
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "went wrong in handler"))
		return
	}

	// active record all day baby
	resp := user.Create()
	u.Respond(c.Writer, resp)
}

func Authenticate(c *gin.Context) {

	user := &models.User{}
	err := json.NewDecoder(c.Request.Body).Decode(user)
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "went wrong in handler"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(c.Writer, resp)
}