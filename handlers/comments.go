// comments.go
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


func CreateComment(c *gin.Context) {
	comment := &models.Comment{}
	err := json.NewDecoder(c.Request.Body).Decode(comment)
	if err != nil {
		u.Respond(c.Writer, u.Message(false, string(err.Error())))
		return
	}

	ok := services.CreateComment(comment)
	if !ok {
		resp := u.Message(false, "could not create comment")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["comment"] = comment
	c.JSON(http.StatusOK, resp)
	return
}

func UpdateComment(c *gin.Context) {
	comment := &models.Comment{}
	err := json.NewDecoder(c.Request.Body).Decode(comment)
	if err != nil {
		u.Respond(c.Writer, u.Message(false, string(err.Error())))
		return
	}

	ok := services.UpdateComment(comment)
	if !ok {
		resp := u.Message(false, "could not update comment")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["comment"] = comment
	c.JSON(http.StatusOK, resp)
	return
}

func DeleteComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "error in your request"))
		return
	}

	ok := services.DeleteComment(id)
	if !ok {
		resp := u.Message(false, "could not delete comment")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := u.Message(true, "success")
	c.JSON(http.StatusOK, resp)
	return
}

func GetComment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		u.Respond(c.Writer, u.Message(false, "error in your request"))
		return
	}

	comment := services.GetComment(id)
	if comment == nil {
		resp := u.Message(false, "could not find comment")
		c.JSON(http.StatusNotFound, resp)
		return
	}
	resp := u.Message(true, "success")
	resp["comment"] = comment
	c.JSON(http.StatusOK, resp)
	return
}

func GetComments(c *gin.Context) {
	comments := services.GetAllComments()
	if len(comments) == 0 {
		resp := u.Message(false, "couldn't get all comments")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := u.Message(true, "success")
	resp["comments"] = comments
	c.JSON(http.StatusOK, resp)
	return
}