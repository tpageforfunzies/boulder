// auth.go
package handlers

import (
	u "github.com/tpageforfunzies/boulder/common"
	"github.com/tpageforfunzies/boulder/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
         "message": "derp",
       },
	)
}


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

	if resp["status"] == false {
		c.JSON(http.StatusForbidden, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func Authenticate(c *gin.Context) {

	user := &models.User{}
	err := json.NewDecoder(c.Request.Body).Decode(user)
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "went wrong in handler"))
		return
	}

	resp := models.Login(user.Email, user.Password)
	if resp["status"] == false {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}