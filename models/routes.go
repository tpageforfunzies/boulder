// routes.go
package models

import (
	u "github.com/tpageforfunzies/boulder/common"
	"github.com/jinzhu/gorm"
	// "fmt"
)

type Route struct {
	gorm.Model
	Name string `json:"name"`
	Grade string `json:"grade"`
	UserId uint `json:"user_id"` //The user that this contact belongs to
}

func (route *Route) Validate() (map[string] interface{}, bool) {

	if route.Name == "" {
		return u.Message(false, "your route needs a name brah"), false
	}

	if route.Grade == "" {
		return u.Message(false, "your route needs a grade my doodie"), false
	}

	if route.UserId <= 0 {
		return u.Message(false, "User not recognized"), false
	}

	return u.Message(true, "success"), true
}

func (route *Route) Create() (map[string] interface{}) {

	check, ok := route.Validate()
	if !ok {
		return check
	}

	GetDB().Create(route)

	resp := u.Message(true, "success")
	resp["route"] = route
	return resp
}

func GetRoute(id int) (*Route) {

	route := &Route{}
	err := GetDB().Table("routes").Where("id = ?", id).First(route).Error
	if err != nil {
		return nil
	}
	return route
}

func GetRoutes(userId int) ([]*Route) {

	routes := make([]*Route, 0)
	err := GetDB().Table("routes").Where("user_id = ?", userId).Find(&routes).Error
	if err != nil {
		return nil
	}

	return routes
}