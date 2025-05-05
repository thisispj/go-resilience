package api

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/api/handlers"
	"microservice/internal/middleware"
	"microservice/pkg/observability"
)

// SetupRouter configures the Gin router with all endpoints and middleware
func SetupRouter(logger observability.Logger, metrics observability.Metrics) *gin.Engine {
	router := gin.New()

	// Add middleware
	router.Use(gin.Recovery())
	router.Use(middleware.LoggingMiddleware(logger))
	router.Use(middleware.MetricsMiddleware(metrics))

	appRoute := router.Group("/app")

	// Health check endpoint
	healthHandler := handlers.NewHealthHandler(logger, metrics)
	appRoute.GET("/health", healthHandler.Check)

	// Metrics endpoint
	appRoute.GET("/metrics", metrics.Handler())

	return router
}
