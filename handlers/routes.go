// routes.go
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


func CreateRoute(c *gin.Context) {
	route := &models.Route{}
	err := json.NewDecoder(c.Request.Body).Decode(route)
	if err != nil {
		u.Respond(c.Writer, u.Message(false, string(err.Error())))
		return
	}

	ok := services.CreateRoute(route)
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

func UpdateRoute(c *gin.Context) {
	route := &models.Route{}
	err := json.NewDecoder(c.Request.Body).Decode(route)
	if err != nil {
		u.Respond(c.Writer, u.Message(false, string(err.Error())))
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
		u.Respond(c.Writer, u.Message(false, "error in your request"))
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
		u.Respond(c.Writer, u.Message(false, "error in your request"))
		return
	}

	route := services.GetRoute(id)
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

func GetRouteComments (c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "error in your request"))
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