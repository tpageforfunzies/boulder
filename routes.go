// routes.go
// Routes stuff here
package boulder

import (
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/tpageforfunzies/boulder/handlers"
	// commented out with log dir
	// "net/http"
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

	// disabling all this while running on kubernetes
	// because os filesystem stuff sucks on k8s

	// Static Content
	// group.Static("/debug", "./debug")

	// Serve the logs directory for the static page
	// only the gin.log file specified open in auth middleware
	// group.StaticFS("/logs", http.Dir("./logs"))

	// User routes
	group.POST("/user/new", handlers.CreateUser)
	group.POST("/user/login", handlers.Authenticate)
	group.GET("/user/:id", handlers.GetUser)
	group.GET("/user/:id/routes", handlers.GetUserRoutes)
	group.GET("/user/:id/comments", handlers.GetUserComments)
	group.GET("/user/:id/followers", handlers.GetUserFollowers)
	group.GET("/user/:id/followed", handlers.GetUserFollowed)
	group.POST("/userpic/:id", handlers.AddProfilePic)

	//Users routes
	group.GET("/users", handlers.GetUsers)

	// Route routes
	group.POST("/route/new", handlers.CreateRoute)
	group.GET("/route/:id", handlers.GetRoute)
	group.PUT("/route/:id", handlers.UpdateRoute)
	group.DELETE("/route/:id", handlers.DeleteRoute)
	group.GET("/route/:id/comments", handlers.GetRouteComments)
	group.POST("/routepic/:id", handlers.AddRoutePic)

	// Routes routes
	group.GET("/routes", handlers.GetRoutes)
	group.GET("/routes/:count", handlers.GetRecentRoutes)

	// Comment Routes
	group.POST("/comment/new", handlers.CreateComment)
	group.GET("/comment/:id", handlers.GetComment)
	group.PUT("/comment/:id", handlers.UpdateComment)
	group.DELETE("/comment/:id", handlers.DeleteComment)

	// Comments Routes
	group.GET("/comments", handlers.GetComments)

	// Relationship Routes
	group.POST("/relationship/new", handlers.CreateRelationship)
	group.GET("/relationship/:id", handlers.GetRelationship)
	group.DELETE("/relationship/:id", handlers.DeleteRelationship)

	// Relationships Routes
	group.GET("/relationships", handlers.GetRelationships)
}
