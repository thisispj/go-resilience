package middleware

import (
	"microservice/pkg/observability"
	"time"

	"github.com/gin-gonic/gin"
)

// MetricsMiddleware collects metrics for each request
func MetricsMiddleware(metrics observability.Metrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// Process request
		c.Next()

		// Record metrics after completion
		path := c.FullPath()
		if path == "" {
			path = "unknown"
		}

		metrics.IncRequestCount(path, c.Request.Method, c.Writer.Status())
		metrics.ObserveRequestDuration(path, time.Since(startTime).Seconds())
	}
}
