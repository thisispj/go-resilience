package middleware

import (
	"microservice/pkg/observability"
	"time"

	"github.com/gin-gonic/gin"
)

// LoggingMiddleware adds structured logging to requests
func LoggingMiddleware(logger observability.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate latency
		latency := time.Since(startTime)

		// Log request details
		logger.Info("Request processed", map[string]interface{}{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"latency_ms": latency.Milliseconds(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		})
	}
}
