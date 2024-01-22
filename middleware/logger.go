package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Process request
		c.Next()

		//get status code
		status := c.Writer.Status()

		if status >= 500 {
			log.Error().
				Int("status", status).
				Str("method", c.Request.Method).
				Time("time", time.Now()).
				Str("path", c.Request.URL.Path).
				Dur("latency", time.Since(start)).
				Msg("Server Error")
		} else if status >= 400 {
			log.Warn().
				Int("status", status).
				Str("method", c.Request.Method).
				Time("time", time.Now()).
				Str("path", c.Request.URL.Path).
				Dur("latency", time.Since(start)).
				Msg("Client Error")
		} else {
			log.Info().
				Int("status", status).
				Str("method", c.Request.Method).
				Time("time", time.Now()).
				Str("path", c.Request.URL.Path).
				Dur("latency", time.Since(start)).
				Msg("Success")
		}
	}
}
