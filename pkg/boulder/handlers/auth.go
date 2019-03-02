// auth.go
package handlers

import (
	"net/http"
	u "github.com/tpageforfunzies/boulder/pkg/boulder/common"
	"github.com/tpageforfunzies/boulder/pkg/boulder/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
)


func CreateUser(c *gin.Context) {

	db := models.GetDB()
	defer db.Close()

	r := c.Request
	w := c.Writer

	user := &models.User{}
	// throw this bitch up in that object
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "went wrong in handler"))
		return
	}

	// active record all day baby
	resp := user.Create()
	u.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		u.Respond(w, u.Message(false, "went wrong in handler"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	u.Respond(w, resp)
}