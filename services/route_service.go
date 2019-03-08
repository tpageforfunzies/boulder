// route_service.go
package services

import (
	u "github.com/tpageforfunzies/boulder/common"
	"github.com/tpageforfunzies/boulder/models"
)

func ValidateRoute(route *models.Route) (map[string] interface{}, bool) {

	if route.Name == "" {
		return u.Message(false, "your route needs a name brah"), false
	}

	if route.Grade == "" {
		return u.Message(false, "your route needs a grade my doodie"), false
	}

	if route.UserId == 0 {
		return u.Message(false, "User not recognized"), false
	}

	return u.Message(true, "success"), true
}

func CreateRoute(route *models.Route) (map[string] interface{}) {

	check, ok := ValidateRoute(route)
	if !ok {
		return check
	}

	GetDB().Create(route)

	resp := u.Message(true, "success")
	resp["route"] = route
	return resp
}

func UpdateRoute(route *models.Route) (map[string] interface{}) {
	check, ok := ValidateRoute(route)
	if !ok {
		return check
	}

	GetDB().Model(&route).Updates(&route)

	resp := u.Message(true, "success")
	resp["route"] = route
	return resp
}

func DeleteRoute(id int) bool {
	err := GetDB().Delete(&models.Route{}, id).Error
	return err == nil
}

func GetRoute(id int) (*models.Route) {
	route := &models.Route{}
	err := GetDB().Table("routes").Where("id = ?", id).First(route).Error
	if err != nil {
		return nil
	}
	return route
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