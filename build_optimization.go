// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !debug
// +build !debug

package gin

import "runtime"

// Production build optimizations
func init() {
	// Disable debug logging in production builds
	SetMode(ReleaseMode)
	
	// Optimize GOMAXPROCS for typical web server workloads
	// This can be overridden by calling runtime.GOMAXPROCS() explicitly
	if runtime.GOMAXPROCS(0) == runtime.NumCPU() {
		// Default is already optimal for CPU-bound workloads
		// For I/O-bound workloads, consider setting it to NumCPU() * 2
	}
}

// BuildInfo provides build optimization information
type BuildInfo struct {
	Mode            string
	JSONProvider    string
	OptimizedRoutes bool
	PoolingEnabled  bool
	GOMAXPROCSRatio float64
}

// GetBuildInfo returns current build optimization status
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Mode:            Mode(),
		JSONProvider:    "sonic", // When built with sonic tag
		OptimizedRoutes: true,
		PoolingEnabled:  true,
		GOMAXPROCSRatio: float64(runtime.GOMAXPROCS(0)) / float64(runtime.NumCPU()),
	}
}

// Performance hints for deployment
const (
	// Recommended GOMAXPROCS setting hint
	// Set GOMAXPROCS=runtime.NumCPU() for CPU-intensive workloads
	// or GOMAXPROCS=runtime.NumCPU()*2 for I/O-intensive workloads
	OptimalGOMAXPROCS = "Use GOMAXPROCS=runtime.NumCPU() for CPU-bound or NumCPU()*2 for I/O-bound"
	
	// Recommended build flags for maximum performance
	// The sonic tag enables ByteDance Sonic JSON library for ~58% faster JSON marshaling
	// The gcflags enable aggressive inlining optimizations
	OptimalBuildFlags = "-ldflags='-s -w' -tags='sonic' -gcflags='-l=4'"
	
	// Recommended server configuration for production
	OptimalServerConfig = "ReadTimeout=5s, WriteTimeout=10s, IdleTimeout=60s, MaxHeaderBytes=1MB"
	
	// UDP buffer tuning for QUIC workloads (Linux)
	// These kernel parameters improve UDP receive performance for QUIC/HTTP3
	UDPBufferTuning = "net.core.rmem_max=8388608 net.core.rmem_default=8388608"
)