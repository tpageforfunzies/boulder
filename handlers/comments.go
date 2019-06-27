// comments.go
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

	comment := services.GetCommentById(id)
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
	countParam := c.DefaultQuery("count", "")
	offsetParam := c.DefaultQuery("offset", "")

	// if we don't have either param
	if countParam == "" || offsetParam == "" {
		comments := services.GetAllComments(0, 0)
		if len(comments) == 0 {
			resp := u.Message(false, "could not find the comments")
			c.JSON(http.StatusNotFound, resp)
			return
		}
		resp := u.Message(true, "success")
		resp["comments"] = comments
		c.JSON(http.StatusOK, resp)
		return
	}

	// if we have both params decode them
	count, err := strconv.Atoi(countParam)
	offset, err := strconv.Atoi(offsetParam)
	if err != nil {
		resp := u.Message(false, "could not decode query parameter(s)")
		fmt.Print(err.Error())
		c.JSON(http.StatusInternalServerError, resp)
		return
	}

	// grab {count} comments start at {offset}
	comments := services.GetAllComments(count, offset)
	if len(comments) == 0 {
		resp := u.Message(false, "could not find the comments")
		c.JSON(http.StatusNotFound, resp)
		return
	}

	resp := u.Message(true, "success")
	resp["comments"] = comments
	c.JSON(http.StatusOK, resp)
	return
}
