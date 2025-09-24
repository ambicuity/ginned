// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build !noprofile
// +build !noprofile

package gin

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnableProfiling(t *testing.T) {
	router := New()
	router.EnableProfiling()

	// Test default pprof endpoint
	w := PerformRequest(router, http.MethodGet, "/debug/pprof/")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestEnableProfilingCustomPath(t *testing.T) {
	router := New()
	router.EnableProfiling("/custom/pprof")

	// Test custom pprof endpoint
	w := PerformRequest(router, http.MethodGet, "/custom/pprof/")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestRuntimeStatsEndpoint(t *testing.T) {
	router := New()
	router.RuntimeStatsEndpoint()

	// Test default stats endpoint
	w := PerformRequest(router, http.MethodGet, "/stats")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	
	// Should contain runtime stats fields
	body := w.Body.String()
	assert.Contains(t, body, "goroutines")
	assert.Contains(t, body, "mem_alloc_bytes")
	assert.Contains(t, body, "gomaxprocs")
}

func TestRuntimeStatsEndpointCustomPath(t *testing.T) {
	router := New()
	router.RuntimeStatsEndpoint("/custom/stats")

	// Test custom stats endpoint
	w := PerformRequest(router, http.MethodGet, "/custom/stats")
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetRuntimeStats(t *testing.T) {
	stats := GetRuntimeStats()
	
	// Verify stats structure
	assert.Greater(t, stats.Goroutines, 0)
	assert.Greater(t, stats.MemAlloc, uint64(0))
	assert.Greater(t, stats.MemSys, uint64(0))
	assert.GreaterOrEqual(t, stats.GCCycles, uint32(0))
	assert.GreaterOrEqual(t, stats.GCPauseTotal, int64(0))
	assert.Greater(t, stats.GOMAXPROCS, 0)
}

func TestNewGCOptimizer(t *testing.T) {
	optimizer := NewGCOptimizer()
	assert.NotNil(t, optimizer)
	assert.Equal(t, 100, optimizer.gcPercent)
}

func TestGCOptimizerSetGCPercent(t *testing.T) {
	optimizer := NewGCOptimizer()
	
	result := optimizer.SetGCPercent(50)
	assert.Equal(t, 50, result)
	assert.Equal(t, 50, optimizer.gcPercent)
}

func TestGCOptimizerOptimizeForLatency(t *testing.T) {
	optimizer := NewGCOptimizer()
	optimizer.OptimizeForLatency()
	assert.Equal(t, 50, optimizer.gcPercent)
}

func TestGCOptimizerOptimizeForThroughput(t *testing.T) {
	optimizer := NewGCOptimizer()
	optimizer.OptimizeForThroughput()
	assert.Equal(t, 200, optimizer.gcPercent)
}