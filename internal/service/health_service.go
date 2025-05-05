package service

import (
	"runtime"
	"time"
)

type HealthService struct {
	startTime time.Time
}

func NewHealthService() *HealthService {
	return &HealthService{
		startTime: time.Now(),
	}
}

func (s *HealthService) CheckHealth() (string, map[string]interface{}) {
	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)

	details := map[string]interface{}{
		"uptime":          time.Since(s.startTime).String(),
		"memory_usage_mb": float64(memStats.Alloc) / 1024 / 1024,
		"goroutines":      runtime.NumGoroutine(),
		"go_version":      runtime.Version(),
		"os":              runtime.GOOS,
		"arch":            runtime.GOARCH,
	}

	return "health", details
}
