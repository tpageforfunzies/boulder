// main.go
// Entry to app
package main

import (
  "fmt"
  "github.com/tpageforfunzies/boulder"
  "github.com/tpageforfunzies/boulder/models"
)
func main() {
  // Load up environmental variables
  models.LoadEnvironment()

  db := models.GetDB()
  defer db.Close()

  // Set the router as the default one shipped with Gin
  router := boulder.GetRouter()

  // Setup route group for the API
  api := router.Group("/v1/")

  // Add routes to route group
  boulder.AddRoutes(api)

  // Start and run the router
  err := router.Run(":80")

  if err != nil {
    fmt.Println("something broke my dude")
    fmt.Println(err)
  }
}