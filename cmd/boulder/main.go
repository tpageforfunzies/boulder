// Entry to app
package main

import (
  "fmt"
  "github.com/tpageforfunzies/boulder/pkg/boulder"
)
func main() {
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