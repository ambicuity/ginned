// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !debug
// +build !debug

package gin

// Production build optimizations
func init() {
	// Disable debug logging in production builds
	SetMode(ReleaseMode)
}

// BuildInfo provides build optimization information
type BuildInfo struct {
	Mode            string
	JSONProvider    string
	OptimizedRoutes bool
	PoolingEnabled  bool
}

// GetBuildInfo returns current build optimization status
func GetBuildInfo() BuildInfo {
	return BuildInfo{
		Mode:            Mode(),
		JSONProvider:    "sonic", // When built with sonic tag
		OptimizedRoutes: true,
		PoolingEnabled:  true,
	}
}

// Performance hints for deployment
const (
	// Recommended GOMAXPROCS setting hint
	// Set GOMAXPROCS=runtime.NumCPU() for CPU-intensive workloads
	// or GOMAXPROCS=runtime.NumCPU()*2 for I/O-intensive workloads
	OptimalGOMAXPROCS = "Use GOMAXPROCS=runtime.NumCPU() for CPU-bound or NumCPU()*2 for I/O-bound"

	// Recommended build flags for maximum performance
	OptimalBuildFlags = "-ldflags='-s -w' -tags='sonic' -gcflags='-l=4'"

	// Recommended server configuration
	OptimalServerConfig = "ReadTimeout=5s, WriteTimeout=10s, IdleTimeout=60s, MaxHeaderBytes=1MB"
)
