// auth.go
package handlers

import (
	u "github.com/tpageforfunzies/boulder/common"
	"github.com/tpageforfunzies/boulder/models"
	"github.com/tpageforfunzies/boulder/services"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	resp := services.CreateUser(user)

	if !resp["status"].(bool) {
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

	resp := services.Login(user.Email, user.Password)
	if !resp["status"].(bool) {
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.JSON(http.StatusOK, resp)
	return
}

func GetUserRoutes(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, u.Message(false, "error in your request"))
		return
	}

	routes := services.GetRoutesByUserId(id)
	if len(routes) == 0 {
		resp := u.Message(false, "could not find their routes")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["routes"] = routes
	c.JSON(http.StatusOK, resp)
	return
}

func GetUserComments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusNotFound, u.Message(false, "error in your request"))
		return
	}

	comments := services.GetCommentsByUserId(id)
	if len(comments) == 0 {
		resp := u.Message(false, "could not find their comments")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["comments"] = comments
	c.JSON(http.StatusOK, resp)
	return
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "error in your request"))
		return
	}

	user := services.GetUserById(id)
	if user == nil {
		resp := u.Message(false, "could not find user")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["user"] = user
	c.JSON(http.StatusOK, resp)
	return

}