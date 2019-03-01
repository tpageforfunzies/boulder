package main
import (
  "net/http"
  "github.com/gin-gonic/gin"
)
func main() {
  // Set the router as the default one shipped with Gin
  router := gin.Default()
  // Setup route group for the API
  api := router.Group("/v1")
  {
    api.GET("/", func(c *gin.Context) {
      c.JSON(http.StatusOK, gin.H {
        "message": "derp",
      })
    })
  }
  // Start and run the server
  router.Run(":1337")
}