// relationship_service.go
package services

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/tpageforfunzies/boulder/models"
)

func ValidateRelationship(relationship *models.Relationship) bool {

	if relationship.FollowedID == 0 {
		return false
	}

	if relationship.FollowerID == 0 {
		return false
	}

	// Set up an object just in case
	check := &models.Relationship{}

	// see if it already exists and if so, put in check object
	err := GetDB().Table("relationships").Where("follower_id = ? AND followed_id = ?", relationship.FollowerID, relationship.FollowedID).First(check).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false
	}

	if check.ID != 0 {
		return false
	}

	return true
}

func CreateRelationship(relationship *models.Relationship) (bool, *models.Relationship) {
	ok := ValidateRelationship(relationship)
	if !ok {
		return false, relationship
	}

	return GetDB().Create(relationship).RowsAffected == 1, relationship
}

func DeleteRelationship(id int) bool {
	return GetDB().Delete(&models.Relationship{}, id).RowsAffected == 1
}

func getRelationshipsByUserId(userId int, relation string, count int, offset int) []*models.Relationship {
	relationships := make([]*models.Relationship, 0)

	if count != 0 {
		err := GetDB().Table("relationships").Where(fmt.Sprintf("%s_id = ?", relation), userId).Limit(count).Offset(offset).Find(&relationships).Error
		if err != nil {
			return nil
		}
		return relationships
	}

	err := GetDB().Table("relationships").Where(fmt.Sprintf("%s_id = ?", relation), userId).Find(&relationships).Error
	if err != nil {
		return nil
	}
	return relationships
}

func GetFollowersByUserId(userId int, count int, offset int) []*models.User {
	// gets relationships where user is the followed one, aka their followers
	relationships := getRelationshipsByUserId(userId, "followed", count, offset)

	followers := make([]*models.User, 0)

	for _, relationship := range relationships {
		// service layer synergy!
		follower := GetUserById(relationship.FollowerID)
		if follower == nil {
			return nil
		}

		// don't need all this jazz
		follower.Routes = nil
		follower.Comments = nil
		followers = append(followers, follower)
	}

	return followers
}

func GetFollowedByUserId(userId int, count int, offset int) []*models.User {
	// gets the relationships where user is the follower, aka their followed users
	relationships := getRelationshipsByUserId(userId, "follower", count, offset)

	followed := make([]*models.User, 0)

	for _, relationship := range relationships {
		followed_user := GetUserById(relationship.FollowedID)
		if followed_user == nil {
			return nil
		}

		followed_user.Routes = nil
		followed_user.Comments = nil
		followed = append(followed, followed_user)
	}

	return followed
}
