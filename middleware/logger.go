// logger.go
// verison of the gin-contrib/logger I modified 
// to log request body data, maybe a PR eventually
package middleware

import (
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"fmt"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Config struct {
	Logger *zerolog.Logger
	// UTC a boolean stating whether to use UTC time zone or local.
	UTC            bool
	SkipPath       []string
	SkipPathRegexp *regexp.Regexp
}

// SetLogger initializes the logging middleware.
func SetLogger(config ...Config) gin.HandlerFunc {
	var newConfig Config
	if len(config) > 0 {
		newConfig = config[0]
	}
	var skip map[string]struct{}
	if length := len(newConfig.SkipPath); length > 0 {
		skip = make(map[string]struct{}, length)
		for _, path := range newConfig.SkipPath {
			skip[path] = struct{}{}
		}
	}

	var sublog zerolog.Logger
	if newConfig.Logger == nil {
		sublog = log.Logger
	} else {
		sublog = *newConfig.Logger
	}

	return func(c *gin.Context) {
		if strings.Contains(c.Request.URL.Path, "favicon") {
			c.Next()
			return
		}

		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		buf := make([]byte, 1024)
		num, _ := c.Request.Body.Read(buf)
		reqBody := string(buf[0:num])
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer([]byte(reqBody)))

		c.Next()
		track := true

		if _, ok := skip[path]; ok {
			track = false
		}

		if track &&
			newConfig.SkipPathRegexp != nil &&
			newConfig.SkipPathRegexp.MatchString(path) {
			track = false
		}

		if track {
			end := time.Now()
			latency := end.Sub(start)
			if newConfig.UTC {
				end = end.UTC()
			}

			// don't log the request bodies on these routes
			privateRoutes := []string{
				"/v1/user/new", 
				"/v1/user/login",
			}
			notLog := privateRoutes
			private := false
			for _, value := range notLog {

				if value == path {
					private = true
				}
			}

			var msg string
			if !private {
				msg = fmt.Sprintf("Request Body: %s", reqBody)
			}
			if len(c.Errors) > 0 {
				msg = c.Errors.String()
			}

			dumplogger := sublog.With().
				Int("status", c.Writer.Status()).
				Str("method", c.Request.Method).
				Str("path", path).
				Str("ip", c.ClientIP()).
				Dur("latency", latency).
				Str("user-agent", c.Request.UserAgent()).
				Logger()

			switch {
			case c.Writer.Status() >= http.StatusBadRequest && c.Writer.Status() < http.StatusInternalServerError:
				{
					dumplogger.Warn().
						Msg(msg)
				}
			case c.Writer.Status() >= http.StatusInternalServerError:
				{
					dumplogger.Error().
						Msg(msg)
				}
			default:
				dumplogger.Info().
					Msg(msg)
			}
		}

	}
}

func AddLogWriters(r *gin.Engine) {
	// Set default to debug for now
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	f, _ := os.Create("logs/gin.log")

	log.Logger = log.Output(
		zerolog.ConsoleWriter{
			Out:     io.MultiWriter(f, os.Stdout),
			NoColor: true,
		},
	)
}