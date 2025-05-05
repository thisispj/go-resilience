package test

import (
	"encoding/json"
	"microservice/internal/api"
	"microservice/pkg/observability"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHealthEndpoint(t *testing.T) {
	// Set Gin to test mode
	gin.SetMode(gin.TestMode)

	// Create dependencies
	logger := observability.NewLogger()
	metrics := observability.NewMetrics()

	// Get router
	router := api.SetupRouter(logger, metrics)

	// Create test recorder and request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/app/health", nil)

	// Perform request
	router.ServeHTTP(w, req)

	// Assert status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Parse response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	// Assert response content
	assert.Nil(t, err)
	assert.Equal(t, "health", response["status"])
	assert.NotNil(t, response["timestamp"])

	// Assert details are present
	details, ok := response["details"].(map[string]interface{})
	assert.True(t, ok)
	assert.NotNil(t, details["uptime"])
	assert.NotNil(t, details["memory_usage_mb"])
	assert.NotNil(t, details["goroutines"])
}
