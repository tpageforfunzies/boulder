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
	"fmt"
	"io"
    "mime"
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

// profile pic functionality right here baby
// gets the post multipart form request from front end
// and gets the id out of the context and goes through each
// file one by one, uploads to s3, and updates user data with the 
// new image url
// does multiple so can have a gallery eventually
func AddProfilePic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := u.Message(false, "need an id for user to update")
		c.JSON(http.StatusBadRequest, resp)
		return
	}
		
	multipart, err := c.Request.MultipartReader()
	if err != nil {
		fmt.Println(err.Error())
	}

	for {
		mimePart, err := multipart.NextPart()
		fmt.Println(mimePart)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		_, params, err := mime.ParseMediaType(mimePart.Header.Get("Content-Disposition"))
		headers := c.Request.Header.Get("Content-Type")
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		// call the image service and upload the parsed file by the filename and type
		ok, imageUrl := services.UploadProfilePicture(params["filename"], headers, mimePart)
		if ok != true {
			// woops if false, imageUrl is the error
			fmt.Println("Errored in S3 service")
			c.JSON(http.StatusInternalServerError, imageUrl)
		}

		// call the user service and update the ImageUrl for User :id with new public url
		ok, _ = services.UpdateUserProfilePic(id, imageUrl)
		if ok != true {
			c.String(http.StatusInternalServerError, "couldn't update user")
		}

		c.String(http.StatusOK, imageUrl)
	}
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