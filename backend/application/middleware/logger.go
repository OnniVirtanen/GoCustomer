package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func RequestLoggerMiddleware(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the IP address and the requested endpoint with a timestamp
		logger.Printf("IP: %s, Method: %s, Endpoint: %s\n", c.ClientIP(), c.Request.Method, c.Request.URL.Path)

		// Process request
		c.Next()
	}
}
