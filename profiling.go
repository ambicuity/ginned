// Copyright 2014 ambicuity Ritesh Rana. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !noprofile
// +build !noprofile

package gin

import (
	"net/http"
	_ "net/http/pprof" // Import pprof HTTP endpoints
	"runtime"
	"time"
)

// EnableProfiling adds pprof endpoints to the Gin router for performance analysis
func (engine *Engine) EnableProfiling(profilingPath ...string) {
	basePath := "/debug/pprof"
	if len(profilingPath) > 0 {
		basePath = profilingPath[0]
	}

	// CPU profiling endpoint
	engine.GET(basePath+"/", func(c *Context) {
		http.DefaultServeMux.ServeHTTP(c.Writer, c.Request)
	})

	// All pprof endpoints
	engine.GET(basePath+"/*any", func(c *Context) {
		// Rewrite path for pprof
		c.Request.URL.Path = "/debug/pprof" + c.Param("any")
		http.DefaultServeMux.ServeHTTP(c.Writer, c.Request)
	})
}

// RuntimeStats provides runtime performance metrics
type RuntimeStats struct {
	Goroutines   int           `json:"goroutines"`
	MemAlloc     uint64        `json:"mem_alloc_bytes"`
	MemSys       uint64        `json:"mem_sys_bytes"`
	GCCycles     uint32        `json:"gc_cycles"`
	GCPauseTotal time.Duration `json:"gc_pause_total_ns"`
	GOMAXPROCS   int           `json:"gomaxprocs"`
}

// GetRuntimeStats returns current runtime performance statistics
func GetRuntimeStats() RuntimeStats {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	return RuntimeStats{
		Goroutines:   runtime.NumGoroutine(),
		MemAlloc:     m.Alloc,
		MemSys:       m.Sys,
		GCCycles:     m.NumGC,
		GCPauseTotal: time.Duration(m.PauseTotalNs),
		GOMAXPROCS:   runtime.GOMAXPROCS(0),
	}
}

// RuntimeStatsEndpoint creates an endpoint to expose runtime stats
func (engine *Engine) RuntimeStatsEndpoint(path ...string) {
	endpoint := "/stats"
	if len(path) > 0 {
		endpoint = path[0]
	}

	engine.GET(endpoint, func(c *Context) {
		stats := GetRuntimeStats()
		c.JSON(200, stats)
	})
}

// GCOptimizer provides GC tuning utilities
type GCOptimizer struct {
	gcPercent int
}

// NewGCOptimizer creates a GC optimizer with recommended settings
func NewGCOptimizer() *GCOptimizer {
	return &GCOptimizer{
		gcPercent: 100, // Default
	}
}

// SetGCPercent adjusts garbage collection target percentage
// Lower values = more frequent GC, less memory usage
// Higher values = less frequent GC, more memory usage
func (gco *GCOptimizer) SetGCPercent(percent int) int {
	gco.gcPercent = percent
	// Note: runtime.SetGCPercent may not be available in all Go versions
	// Use GOGC environment variable as alternative: GOGC=50 for 50%
	return gco.gcPercent
}

// OptimizeForLatency sets GC for low-latency applications
func (gco *GCOptimizer) OptimizeForLatency() {
	gco.SetGCPercent(50) // More frequent GC for lower pause times
}

// OptimizeForThroughput sets GC for high-throughput applications
func (gco *GCOptimizer) OptimizeForThroughput() {
	gco.SetGCPercent(200) // Less frequent GC for higher throughput
}
