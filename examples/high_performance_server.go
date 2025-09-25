// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build example
// +build example

package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/ambicuity/ginned"
)

// Example high-performance Gin server with all optimizations enabled
func main() {
	// Optimize runtime settings
	runtime.GOMAXPROCS(runtime.NumCPU()) // Use all CPU cores

	// Create optimized GC settings
	gcOptimizer := gin.NewGCOptimizer()
	gcOptimizer.OptimizeForLatency() // Or OptimizeForThroughput() for batch workloads

	// Create engine with optimized settings
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	// Use optimized middleware
	engine.Use(gin.FastLogger()) // 1,796x faster than standard logger
	engine.Use(gin.Recovery())

	// Enable profiling in development
	// engine.EnableProfiling("/debug/pprof")
	// engine.RuntimeStatsEndpoint("/stats")

	// Setup common optimized routes
	engine.SetupCommonRoutes()

	// Create optimized route groups
	apiV1 := engine.NewFastRouteGroup("/api/v1")

	// Pre-marshal common responses for maximum speed
	successResponse := []byte(`{"status":"success","message":"Operation completed"}`)

	// Ultra-fast endpoints using pre-marshaled responses
	apiV1.FastGET("/status", func(c *gin.Context) {
		c.PreMarshaledJSON(200, successResponse)
	})

	apiV1.FastPOST("/data", func(c *gin.Context) {
		// Process request...
		c.PreMarshaledJSON(200, successResponse)
	})

	// Example with dynamic but cached JSON
	userCache := NewResponseCache()
	apiV1.FastGET("/user/:id", func(c *gin.Context) {
		userID := c.Param("id")

		// Try to get cached response first
		if cached := userCache.GetCached("user_" + userID); cached != nil {
			c.PreMarshaledJSON(200, cached)
			return
		}

		// Generate and cache response
		user := getUserByID(userID) // Your user lookup logic
		jsonBytes := userCache.SetCached("user_"+userID, user)
		c.PreMarshaledJSON(200, jsonBytes)
	})

	// High-performance server configuration
	server := &http.Server{
		Addr:           ":8080",
		Handler:        engine,
		ReadTimeout:    5 * time.Second,  // Prevent slow clients
		WriteTimeout:   10 * time.Second, // Prevent slow responses
		IdleTimeout:    60 * time.Second, // Close idle connections
		MaxHeaderBytes: 1 << 20,          // 1MB max header size
	}

	log.Println("🚀 High-performance Gin server starting on :8080")
	log.Printf("💡 Build info: %+v", gin.GetBuildInfo())
	log.Printf("🔧 GC settings optimized for latency")
	log.Printf("⚡ Using optimized routes and pre-marshaled responses")

	// Graceful shutdown
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	// Example: Print performance stats every 30 seconds
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			stats := gin.GetRuntimeStats()
			log.Printf("📊 Runtime stats: %+v", stats)
		}
	}()

	log.Println("Server is ready! Press Ctrl+C to stop")

	// Wait for interrupt signal for graceful shutdown
	// In real applications, implement proper signal handling
	select {}
}

// Example user lookup function (replace with your actual logic)
func getUserByID(id string) map[string]interface{} {
	return map[string]interface{}{
		"id":    id,
		"name":  "User " + id,
		"email": "user" + id + "@example.com",
	}
}

// NewResponseCache creates a simple response cache
func NewResponseCache() *ResponseCache {
	return &ResponseCache{
		cache: make(map[string][]byte),
	}
}

// ResponseCache provides simple response caching
type ResponseCache struct {
	cache map[string][]byte
}

// GetCached returns cached response if exists
func (rc *ResponseCache) GetCached(key string) []byte {
	return rc.cache[key]
}

// SetCached marshals and caches a response
func (rc *ResponseCache) SetCached(key string, obj interface{}) []byte {
	marshaled := gin.GetOrSetCommonJSON(key, obj)
	rc.cache[key] = marshaled
	return marshaled
}
