// route_service.go
package services

import (
	"github.com/tpageforfunzies/boulder/models"
)

func ValidateRoute(route *models.Route) bool {

	if route.Name == "" {
		return false
	}

	if route.Grade == "" {
		return false
	}

	if route.UserId == 0 {
		return false
	}

	return true
}

func CreateRoute(route *models.Route) (bool, *models.Route) {
	ok := ValidateRoute(route)
	if !ok {
		return false, route
	}

	return GetDB().Create(route).RowsAffected == 1, route
}

func UpdateRoute(route *models.Route) bool {
	ok := ValidateRoute(route)
	if !ok {
		return false
	}

	return GetDB().Model(&route).Updates(&route).RowsAffected == 1
}

func DeleteRoute(id int) bool {
	return GetDB().Delete(&models.Route{}, id).RowsAffected == 1
}

func GetRouteById(id int) *models.Route {
	route := &models.Route{}
	err := GetDB().Table("routes").Preload("Comments").Find(route, id).Error
	if err != nil {
		return nil
	}

	return route
}

func GetAllRoutes(count int, offset int) []*models.Route {
	routes := make([]*models.Route, 0)
	if count != 0 {
		err := GetDB().Preload("Comments").Limit(count).Offset(offset).Order("created_at desc", true).Find(&routes).Error
		if err != nil {
			return nil
		}
		return routes
	}
	err := GetDB().Preload("Comments").Find(&routes).Error
	if err != nil {
		return nil
	}
	return routes
}

func GetRoutesByUserId(userId int, count int, offset int) []*models.Route {
	routes := make([]*models.Route, 0)
	if count != 0 {
		err := GetDB().Table("routes").Where("user_id = ?", userId).Preload("Comments").Limit(count).Offset(offset).Order("created_at desc", true).Find(&routes).Error
		if err != nil {
			return nil
		}
		return routes
	}
	err := GetDB().Table("routes").Where("user_id = ?", userId).Preload("Comments").Find(&routes).Error
	if err != nil {
		return nil
	}
	return routes
}

func UpdateRoutePic(id int, imageUrl string) (bool, string) {
	db := GetDB()
	route := GetRouteById(id)
	route.ImageUrl = imageUrl
	return db.Save(&route).RowsAffected == 1, imageUrl
}

func GetRecentRoutes(count int) []*models.Route {
	routes := make([]*models.Route, count)
	err := GetDB().Table("routes").Order("created_at desc").Limit(count).Find(&routes).Error
	if err != nil {
		return nil
	}
	return routes
}
