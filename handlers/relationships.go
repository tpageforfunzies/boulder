// relationships.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	u "github.com/tpageforfunzies/boulder/common"
	"github.com/tpageforfunzies/boulder/models"
	"github.com/tpageforfunzies/boulder/services"
)

func CreateRelationship(c *gin.Context) {
	relationship := &models.Relationship{}
	err := json.NewDecoder(c.Request.Body).Decode(relationship)

	if err != nil {
		fmt.Println("here")
		resp := u.Message(false, string(err.Error()))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ok, relationship := services.CreateRelationship(relationship)
	if !ok {
		resp := u.Message(false, "could not create relationship")
		resp["relationship"] = relationship
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["relationship"] = relationship
	c.JSON(http.StatusOK, resp)
	return
}

func DeleteRelationship(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resp := u.Message(false, string(err.Error()))
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	ok := services.DeleteRelationship(id)
	if !ok {
		resp := u.Message(false, "could not delete relationship")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := u.Message(true, "success")
	c.JSON(http.StatusOK, resp)
	return
}
