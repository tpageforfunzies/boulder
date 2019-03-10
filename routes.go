// routes.go
// Routes stuff here
package boulder

import (
	"github.com/tpageforfunzies/boulder/handlers"
	"github.com/gin-gonic/gin"
	"sync"
)

var router *gin.Engine
var once sync.Once

// GetRouter returns a singleton instance of the router
func GetRouter() *gin.Engine {
	// Singleton router instance
	once.Do(func() {
		router = gin.Default()
	})

	return router
}

func AddApiRoutes(group *gin.RouterGroup) {
	group.GET("/", handlers.HomeHandler)

	// Static Content
	group.Static("/content", "./content")

	// User routes
	group.POST("/user/new", handlers.CreateUser)
	group.POST("/user/login", handlers.Authenticate)
	// group.GET("/user/routes", )
	// group.GET("/user/comments", )

	//Users routes
	group.GET("/users", handlers.GetUsers)

	group.GET("/user/:id", handlers.GetUser)
	group.GET("/user/:id/routes", handlers.GetUserRoutes)
	group.GET("/user/:id/comments", handlers.GetUserComments)

	// Route routes
	group.POST("/route/new", handlers.CreateRoute)
	group.GET("/route/:id", handlers.GetRoute)
	group.PUT("/route/:id", handlers.UpdateRoute)
	group.DELETE("/route/:id", handlers.DeleteRoute)
	group.GET("/route/:id/comments", handlers.GetRouteComments)

	// Routes routes
	group.GET("/routes", handlers.GetRoutes)

	// Comment Routes
	group.POST("/comment/new", handlers.CreateComment)
	group.GET("/comment/:id", handlers.GetComment)
	group.PUT("/comment/:id", handlers.UpdateComment)
	group.DELETE("/comment/:id", handlers.DeleteComment)

	// Comments Routes
	group.GET("/comments", handlers.GetComments)
}