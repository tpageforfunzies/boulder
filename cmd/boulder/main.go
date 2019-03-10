// main.go
// Entry to app
package main

import (
  "fmt"
  "github.com/tpageforfunzies/boulder"
  "github.com/tpageforfunzies/boulder/services"
  "github.com/tpageforfunzies/boulder/middleware"
)

func main() {
  // Load up environmental variables
  services.LoadEnvironment()

  db := services.GetDB()
  defer db.Close()

  // Set the router as the default one shipped with Gin
  router := boulder.GetRouter()

  // Add log writers to router and 
  // add logging and auth middleware
  middleware.AddLogWriters(router)
  router.Use(middleware.SetLogger())
  router.Use(middleware.JwtAuthentication)
  
  // Setup route group for the API
  api := router.Group("/v1/")

  // Add routes to route group
  boulder.AddApiRoutes(api)

  // Start and run the router
  err := router.Run(":80")

  if err != nil {
    fmt.Println("something broke my dude")
    fmt.Println(err)
  }
}