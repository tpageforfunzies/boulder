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
	"strings"
)

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
         "message": "derp",
       },
	)
}

func CreateUser(c *gin.Context) {

	user := &models.User{}
	err := json.NewDecoder(c.Request.Body).Decode(user)
	if err != nil {
		resp := u.Message(false, "went wrong in handler")
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result, user := services.CreateUser(user)

	// make this less sketchy
	if result != "User created!" {
		resp := u.Message(false, result)
		c.JSON(http.StatusForbidden, resp)
		return
	}

	resp := u.Message(true, result)
	resp["user"] = user
	c.JSON(http.StatusOK, resp)
	return
}

func Authenticate(c *gin.Context) {

	user := &models.User{}
	err := json.NewDecoder(c.Request.Body).Decode(user)
	if err != nil {
		resp := u.Message(false, "went wrong in handler")
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result, user := services.Login(strings.ToLower(user.Email), user.Password)
	if result != "Logged In" {
		resp := u.Message(false, result)
		c.JSON(http.StatusForbidden, resp)
		return
	}
	
	resp := u.Message(true, result)
	resp["user"] = user
	c.JSON(http.StatusOK, resp)
	return
}

func GetUserRoutes(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := u.Message(false, "went wrong in handler")
		c.JSON(http.StatusBadRequest, resp)
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
		resp := u.Message(false, "went wrong in handler")
		c.JSON(http.StatusBadRequest, resp)
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
		resp := u.Message(false, "went wrong in handler")
		c.JSON(http.StatusBadRequest, resp)
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

func GetUsers(c *gin.Context) {
	users := services.GetAllUsers()
	if len(users) == 0 {
		resp := u.Message(false, "couldn't get all users")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := u.Message(true, "success")
	resp["users"] = users
	c.JSON(http.StatusOK, resp)
	return
}