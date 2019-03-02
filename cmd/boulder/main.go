// Entry to app
package main

import (
  "fmt"
  "github.com/tpageforfunzies/boulder/pkg/boulder"
  "github.com/tpageforfunzies/boulder/pkg/boulder/models"
  "log"
)
func main() {
  // Load up environmental variables
  models.LoadEnvironment()
  log.Output(1, "loaded env")
  db := models.GetDB()
  log.Output(1, "got db")
  defer db.Close()
  // if err != nil {
  //   panic(err)
  // }

  // Set the router as the default one shipped with Gin
  router := boulder.GetRouter()

  // Setup route group for the API
  api := router.Group("/v1/")

  // Add routes to route group
  boulder.AddRoutes(api)

  // Start and run the router
  err := router.Run(":420")

  if err != nil {
    fmt.Println("something broke my dude")
    fmt.Println(err)
  }
}