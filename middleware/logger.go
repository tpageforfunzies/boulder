// logger.go
package middleware

import (
	"os"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"io"
)

func AddLogger(r *gin.Engine) {
	// Set default to debug for now
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	f, _ := os.Create("gin.log")

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     io.MultiWriter(f, os.Stdout),
			NoColor: false,
		},
	)

	// Add the logger middleware
	r.Use(logger.SetLogger())
}