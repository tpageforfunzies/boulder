// routes.go
package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	u "github.com/tpageforfunzies/boulder/common"
	"github.com/tpageforfunzies/boulder/models"
	"github.com/tpageforfunzies/boulder/services"
)

func CreateRoute(c *gin.Context) {
	route := &models.Route{}
	err := json.NewDecoder(c.Request.Body).Decode(route)

	if err != nil {
		resp := u.Message(false, string(err.Error()))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ok, route := services.CreateRoute(route)
	if !ok {
		resp := u.Message(false, "could not create route")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["route"] = route
	c.JSON(http.StatusOK, resp)
	return
}

// route pic functionality right here baby
// gets the post multipart form request from front end
// and gets the id out of the context and goes through each
// file one by one, uploads to s3, and updates route data with the
// new image url
// does multiple so can have a gallery eventually
func AddRoutePic(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := u.Message(false, "need an id for route to update")
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
		_, _, err = mime.ParseMediaType(mimePart.Header.Get("Content-Disposition"))
		headers := c.Request.Header.Get("Content-Type")
		if err != nil {
			fmt.Println(err.Error())
			break
		}

		fileName := fmt.Sprintf("route%d.jpeg", id)
		// call the image service and upload the parsed file by the filename and type
		ok, imageUrl := services.UploadPicture(fileName, headers, mimePart)
		if !ok {
			// woops if false, imageUrl is the error
			fmt.Println("Errored in S3 service")
			c.JSON(http.StatusInternalServerError, imageUrl)
		}

		// call the route service and update the ImageUrl for route :id with new public url
		ok, _ = services.UpdateRoutePic(id, imageUrl)
		if !ok {
			c.String(http.StatusInternalServerError, "couldn't update route")
		}

		c.String(http.StatusOK, imageUrl)
	}
}

func UpdateRoute(c *gin.Context) {
	route := &models.Route{}
	err := json.NewDecoder(c.Request.Body).Decode(route)
	if err != nil {
		resp := u.Message(false, string(err.Error()))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ok := services.UpdateRoute(route)
	if !ok {
		resp := u.Message(false, "could not update route")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["route"] = route
	c.JSON(http.StatusOK, resp)
	return
}

func DeleteRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := u.Message(false, string(err.Error()))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ok := services.DeleteRoute(id)
	if !ok {
		resp := u.Message(false, "could not delete route")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := u.Message(true, "success")
	c.JSON(http.StatusOK, resp)
	return
}

func GetRoute(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := u.Message(false, string(err.Error()))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	route := services.GetRouteById(id)
	if route == nil {
		resp := u.Message(false, "could not find route")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["route"] = route
	c.JSON(http.StatusOK, resp)
	return
}

func GetRoutes(c *gin.Context) {
	routes := services.GetAllRoutes()
	if len(routes) == 0 {
		resp := u.Message(false, "couldn't get all routes")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := u.Message(true, "success")
	resp["routes"] = routes
	c.JSON(http.StatusOK, resp)
	return
}

func GetRecentRoutes(c *gin.Context) {
	count, err := strconv.Atoi(c.Param("count"))
	if err != nil {
		resp := u.Message(false, string(err.Error()))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	routes := services.GetRecentRoutes(count)
	if len(routes) == 0 {
		resp := u.Message(false, "couldn't get the recent routes")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := u.Message(true, "success")
	resp["routes"] = routes
	c.JSON(http.StatusOK, resp)
	return
}

func GetRouteComments(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := u.Message(false, string(err.Error()))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	comments := services.GetCommentsByRouteId(id)
	if comments == nil {
		resp := u.Message(false, "couldn't comments for that route")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := u.Message(true, "success")
	resp["comments"] = comments
	c.JSON(http.StatusOK, resp)
	return
}
