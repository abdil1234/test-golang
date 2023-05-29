package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sbecker/gin-api-demo/util"
	log "github.com/sirupsen/logrus"
)

func LogFormatter() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process Request
		c.Next()

		// Stop timer
		duration := util.GetDurationInMillseconds(start)

		entry := log.WithFields(log.Fields{
			"client_ip":  util.GetClientIP(c),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("request_id"),
		})
		entry.Logger.SetFormatter(&log.JSONFormatter{})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
