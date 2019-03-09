// comment_service.go
package services

import (
	"github.com/tpageforfunzies/boulder/models"
)

func ValidateComment(comment *models.Comment) bool {

	if comment.RouteId == 0 {
		return false
	}

	if comment.Content == "" {
		return false
	}

	if comment.UserId == 0 {
		return false
	}

	return true
}

func CreateComment(comment *models.Comment) bool {

	ok := ValidateComment(comment)
	if !ok {
		return false
	}

	return GetDB().Create(comment).RowsAffected == 1
}

func UpdateComment(comment *models.Comment) bool {
	ok := ValidateComment(comment)
	if !ok {
		return false
	}

	return GetDB().Model(&comment).Updates(&comment).RowsAffected == 1
}

func DeleteComment(id int) bool {
	damage := GetDB().Delete(&models.Comment{}, id).RowsAffected
	return damage == 1
}

func GetComment(id int) (*models.Comment) {
	comment := &models.Comment{}
	err := GetDB().Table("comments").Where("id = ?", id).First(comment).Error
	if err != nil {
		return nil
	}
	return comment
}

func GetAllComments() ([]*models.Comment) {
	comments := make([]*models.Comment, 0)
	err := GetDB().Find(&comments).Error
	if err != nil {
		return nil
	}
	return comments
}

func GetCommentsByUserId(userId int) ([]*models.Comment) {
	comments := make([]*models.Comment, 0)
	err := GetDB().Table("comments").Where("user_id = ?", userId).Find(&comments).Error
	if err != nil {
		return nil
	}

	return comments
}

func GetCommentsByRouteId(id int) ([]*models.Comment) {
	comments := make([]*models.Comment, 0)
	err := GetDB().Table("comments").Where("route_id = ?", id).Find(&comments).Error

	if err != nil {
		return nil
	}

	return comments
}