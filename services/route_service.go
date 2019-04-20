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

func GetRouteById(id int) (*models.Route) {
	route := &models.Route{}
	return FindSingleById(route, id).(*models.Route)
}

func GetAllRoutes() ([]*models.Route) {
	routes := make([]*models.Route, 0)
	err := GetDB().Find(&routes).Error
	if err != nil {
		return nil
	}
	return routes
}

func GetRoutesByUserId(userId int) ([]*models.Route) {
	routes := make([]*models.Route, 0)
	 err := GetDB().Table("routes").Where("user_id = ?", userId).Find(&routes).Error
       if err != nil {
               return nil
       }
	return routes
}