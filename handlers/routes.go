// routes.go
package handlers

import (
	u "github.com/tpageforfunzies/boulder/common"
	"github.com/tpageforfunzies/boulder/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)


func CreateRoute(c *gin.Context) {

	route := &models.Route{}
	// throw this bitch up in that object
	err := json.NewDecoder(c.Request.Body).Decode(route)
	if err != nil {
		u.Respond(c.Writer, u.Message(false, string(err.Error())))
		return
	}

	// active record all day baby
	resp := route.Create()
	u.Respond(c.Writer, resp)
}

func GetRoutesFor(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "error in your request"))
		return
	}

	data := models.GetRoutes(id)
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(c.Writer, resp)
}