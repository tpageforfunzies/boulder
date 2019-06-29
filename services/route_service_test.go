package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tpageforfunzies/boulder/models"
)

func TestValidateRouteSuccess(t *testing.T) {
	routeOne := &models.Route{}
	routeOne.Name = "testerson"
	routeOne.Grade = "a million"
	routeOne.UserId = 1
	assert.True(t, ValidateRoute(routeOne), "should be true")
}

func TestValidateRouteNoName(t *testing.T) {
	routeOne := &models.Route{}
	routeOne.Name = ""
	routeOne.Grade = "a million"
	routeOne.UserId = 1
	assert.False(t, ValidateRoute(routeOne), "should be false, missing name")
}

func TestValidateRouteNoGrade(t *testing.T) {
	routeOne := &models.Route{}
	routeOne.Name = "testyboi"
	routeOne.Grade = ""
	routeOne.UserId = 1
	assert.False(t, ValidateRoute(routeOne), "should be false, missing grade")
}

func TestValidateRouteNoUserId(t *testing.T) {
	routeOne := &models.Route{}
	routeOne.Name = "testyboi"
	routeOne.Grade = "at least 2"
	routeOne.UserId = 0
	assert.False(t, ValidateRoute(routeOne), "should be false, missing user id")
}
