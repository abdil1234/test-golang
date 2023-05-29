package middlewares

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const REQUEST_ID_KEY = "request_id"

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.Request.Header.Get(REQUEST_ID_KEY)
		if requestID == "" {
			requestID = fmt.Sprintf("%s", uuid.NewV4())
		}

		ctx := context.WithValue(c.Request.Context(), REQUEST_ID_KEY, requestID)
		c.Request = c.Request.WithContext(ctx)

		c.Set(REQUEST_ID_KEY, requestID)
		c.Writer.Header().Set(REQUEST_ID_KEY, requestID)
		c.Next()
	}
}
