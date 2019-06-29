// comment_service_test.go
package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpageforfunzies/boulder/models"
)

func TestValidateCommentSuccess(t *testing.T) {
	comment := &models.Comment{}
	comment.RouteId = 1
	comment.Content = "a million"
	comment.UserId = 1
	assert.True(t, ValidateComment(comment), "should be true")
}

func TestValidateCommentRouteId(t *testing.T) {
	comment := &models.Comment{}
	comment.RouteId = 0
	comment.Content = "a million"
	comment.UserId = 1
	assert.False(t, ValidateComment(comment), "should be false, missing routeid")
}

func TestValidateCommentNoContent(t *testing.T) {
	comment := &models.Comment{}
	comment.RouteId = 1
	comment.Content = ""
	comment.UserId = 1
	assert.False(t, ValidateComment(comment), "should be false, missing content")
}

func TestValidateCommentNoUserId(t *testing.T) {
	comment := &models.Comment{}
	comment.RouteId = 1
	comment.Content = "at least 2"
	comment.UserId = 0
	assert.False(t, ValidateComment(comment), "should be false, missing user id")
}
