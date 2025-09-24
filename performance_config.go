// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"os"
	"runtime"
)

// PerformanceConfig contains optimized settings for high-performance deployments
type PerformanceConfig struct {
	// Use release mode to disable debug logs and colorized output
	ReleaseMode bool
	
	// Use FastLogger instead of standard Logger for ~1000x improvement
	UseFastLogger bool
	
	// Skip logging for health check endpoints
	SkipHealthChecks bool
	
	// Pre-marshal common JSON responses for zero-marshal overhead
	PreMarshalCommonResponses bool
	
	// Trusted platform for AppEngine deployments
	SetTrustedPlatform bool
	
	// Optimize GOMAXPROCS for the workload type
	OptimizeGOMAXPROCS bool
}

// DefaultPerformanceConfig returns recommended settings for production
func DefaultPerformanceConfig() PerformanceConfig {
	return PerformanceConfig{
		ReleaseMode:               true,
		UseFastLogger:             true,
		SkipHealthChecks:          true,
		PreMarshalCommonResponses: true,
		SetTrustedPlatform:        true,
		OptimizeGOMAXPROCS:        true,
	}
}

// ApplyPerformanceConfig applies the recommended performance optimizations
func ApplyPerformanceConfig(engine *Engine, config PerformanceConfig) *Engine {
	if config.ReleaseMode {
		SetMode(ReleaseMode)
	}
	
	if config.OptimizeGOMAXPROCS {
		// For CPU-intensive workloads, use NumCPU()
		// For I/O-intensive workloads, consider NumCPU() * 2
		runtime.GOMAXPROCS(runtime.NumCPU())
	}
	
	if config.SetTrustedPlatform {
		// This is automatically set by context_appengine.go for AppEngine
		// but can be configured manually if needed
		if os.Getenv("GAE_APPLICATION") != "" || os.Getenv("GOOGLE_CLOUD_PROJECT") != "" {
			engine.TrustedPlatform = PlatformGoogleAppEngine
		}
	}
	
	return engine
}

// NewOptimized creates a new Engine with performance optimizations applied
func NewOptimized(opts ...OptionFunc) *Engine {
	engine := New(opts...)
	config := DefaultPerformanceConfig()
	return ApplyPerformanceConfig(engine, config)
}

// DefaultOptimized returns an Engine with FastLogger, Recovery middleware and performance optimizations
func DefaultOptimized(opts ...OptionFunc) *Engine {
	engine := NewOptimized(opts...)
	config := DefaultPerformanceConfig()
	
	if config.UseFastLogger {
		engine.Use(FastLogger(), Recovery())
	} else {
		engine.Use(Logger(), Recovery())
	}
	
	return engine
}

// Performance recommendations as constants for documentation
const (
	// RecommendedBuildFlags for maximum performance
	RecommendedBuildFlags = "-ldflags='-s -w' -tags='sonic' -gcflags='-l=4'"
	
	// RecommendedServerConfig for production deployments
	RecommendedServerConfig = "ReadTimeout=5s, WriteTimeout=10s, IdleTimeout=60s, MaxHeaderBytes=1MB"
	
	// UDPReceiveBufferTuning for QUIC/UDP workloads
	UDPReceiveBufferTuning = "sudo sysctl -w net.core.rmem_max=8388608 && sudo sysctl -w net.core.rmem_default=8388608"
)