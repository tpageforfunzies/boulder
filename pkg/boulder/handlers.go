// Handlers
package boulder

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func homeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H {
         "message": "derp",
       },
	)
}