// auth.go
package middleware

import (
	"net/http"
	u "github.com/tpageforfunzies/boulder/common"
	"strings"
	"github.com/tpageforfunzies/boulder/models"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"context"
	"github.com/gin-gonic/gin"
)

func JwtAuthentication(c *gin.Context) {
	r := c.Request
	w := c.Writer

	openRoutes := []string{
		// more filesystem stuff commented out
		// while running on k8s
		// "/v1/debug/log.html", 
		// "/v1/debug/js/log.js",
		// "/v1/logs/gin.log",

		// cant auth on login/new user
		"/v1/user/new", 
		"/v1/user/login",
		
		// this is for the front page display
		"/v1/routes/10",
	}

	notAuth := openRoutes
	requestPath := r.URL.Path // current request path

	// check if request does not need authentication, serve the request if it doesn't need it
	for _, value := range notAuth {

		if value == requestPath {
			c.Next()
			return
		}
	}

	tokenHeader := r.Header.Get("Authorization") // Grab the token from the header

	if tokenHeader == "" { // Token is missing, returns with error code 403 Unauthorized
		response := u.Message(false, "Missing auth token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		c.Abort()
		return
	}

	splitted := strings.Split(tokenHeader, " ") // The token normally comes in format `Bearer {token-body}`
	if len(splitted) != 2 {
		response := u.Message(false, "Invalid/Malformed auth token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		c.Abort()
		return
	}

	tokenPart := splitted[1] // Grab the token part, that good good
	tk := &models.Token{}

	token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	})

	if err != nil { // Malformed token, returns with http code 403
		response := u.Message(false, "Malformed authentication token")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		c.Abort()
		return
	}

	if !token.Valid { // you played yourself
		response := u.Message(false, "Token is not valid.")
		w.WriteHeader(http.StatusForbidden)
		w.Header().Add("Content-Type", "application/json")
		u.Respond(w, response)
		c.Abort()
		return
	}

	// Everything went well, proceed with the request and set the caller to the user retrieved from the parsed token
	ctx := context.WithValue(r.Context(), "user", tk.UserId)
	r = r.WithContext(ctx)
	c.Next() // proceed in the middleware chain
}
