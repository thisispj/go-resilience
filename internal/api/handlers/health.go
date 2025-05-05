package handlers

import (
	"microservice/internal/service"
	"microservice/pkg/observability"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	logger        observability.Logger
	metrics       observability.Metrics
	healthService *service.HealthService
}

// NewHealthHandler creates a new health check handler
func NewHealthHandler(logger observability.Logger, metrics observability.Metrics) *HealthHandler {
	return &HealthHandler{
		logger:        logger,
		metrics:       metrics,
		healthService: service.NewHealthService(),
	}
}

// Check handles the health check endpoint
func (h *HealthHandler) Check(c *gin.Context) {
	startTime := time.Now()
	defer func() {
		h.metrics.ObserveRequestDuration("health_check", time.Since(startTime).Seconds())
	}()

	status, healthInfo := h.healthService.CheckHealth()

	h.logger.Info("Health check performed", map[string]interface{}{
		"status":  status,
		"details": healthInfo,
	})

	c.JSON(http.StatusOK, gin.H{
		"status":    status,
		"timestamp": time.Now(),
		"details":   healthInfo,
	})
}
