package observability

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics interface for application metrics
type Metrics interface {
	IncRequestCount(path, method string, statusCode int)
	ObserveRequestDuration(path string, durationSeconds float64)
	Handler() gin.HandlerFunc
}

// PrometheusMetrics implements Metrics using Prometheus
type PrometheusMetrics struct {
	requestCount    *prometheus.CounterVec
	requestDuration *prometheus.HistogramVec
}

// NewMetrics creates a new metrics collector
func NewMetrics() Metrics {
	requestCount := promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path", "method", "status"},
	)

	requestDuration := promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path"},
	)

	return &PrometheusMetrics{
		requestCount:    requestCount,
		requestDuration: requestDuration,
	}
}

// IncRequestCount increments the request counter for a path
func (m *PrometheusMetrics) IncRequestCount(path, method string, statusCode int) {
	m.requestCount.WithLabelValues(path, method, http.StatusText(statusCode)).Inc()
}

// ObserveRequestDuration records the duration of a request
func (m *PrometheusMetrics) ObserveRequestDuration(path string, durationSeconds float64) {
	m.requestDuration.WithLabelValues(path).Observe(durationSeconds)
}

// Handler returns a handler for the metrics endpoint
func (m *PrometheusMetrics) Handler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
