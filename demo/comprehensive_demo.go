// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

//go:build demo
// +build demo

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ComprehensiveDemo demonstrates all the optimized features of the Ginned framework
func main() {
	fmt.Println("🚀 Gin Framework - Comprehensive Optimization Demo")
	fmt.Println("=================================================")

	// 1. Build Optimization Info
	buildInfo := gin.GetBuildInfo()
	fmt.Printf("📊 Build Info: %+v\n", buildInfo)

	// 2. GC Optimization
	gcOptimizer := gin.NewGCOptimizer()
	gcOptimizer.OptimizeForLatency()
	fmt.Println("🗑️  GC optimized for latency")

	// 3. Create engine with optimizations
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// 4. Use optimized middleware
	router.Use(gin.FastLogger())
	router.Use(gin.Recovery())
	fmt.Println("⚡ Using FastLogger middleware")

	// 5. Setup common optimized routes
	router.SetupCommonRoutes()
	fmt.Println("🛣️  Common routes setup: /health, /ping, /version")

	// 6. Profiling endpoints (optional)
	router.EnableProfiling()
	router.RuntimeStatsEndpoint()
	fmt.Println("🔍 Profiling enabled: /debug/pprof/*, /stats")

	// 7. Fast route groups
	apiV1 := router.NewFastRouteGroup("/api/v1")
	fmt.Println("🏗️  Fast route group created: /api/v1")

	// 8. Fast response methods
	apiV1.FastGET("/fast-ok", func(c *gin.Context) {
		c.FastOk()
	})

	apiV1.FastGET("/fast-success", func(c *gin.Context) {
		c.FastSuccess()
	})

	apiV1.FastGET("/fast-pong", func(c *gin.Context) {
		c.FastPong()
	})

	apiV1.FastGET("/fast-error", func(c *gin.Context) {
		c.FastError()
	})

	apiV1.FastGET("/fast-number", func(c *gin.Context) {
		c.FastJSONNumber(200, 42)
	})

	// 9. Optimized context methods
	apiV1.FastGET("/fast-json", func(c *gin.Context) {
		data := map[string]interface{}{
			"message": "This uses FastJSON",
			"time":    time.Now(),
			"status":  "optimized",
		}
		c.FastJSON(200, data)
	})

	apiV1.FastGET("/string-fast", func(c *gin.Context) {
		c.StringFast(200, "This is optimized string response!")
	})

	// 10. Pre-marshaled responses for ultimate speed
	successResponse := []byte(`{"status":"success","optimization":"pre-marshaled","speed":"maximum"}`)
	apiV1.FastGET("/pre-marshaled", func(c *gin.Context) {
		c.PreMarshaledJSON(200, successResponse)
	})

	// 11. Cached responses
	userCache := map[string][]byte{
		"123": gin.GetOrSetCommonJSON("user_123", map[string]interface{}{
			"id":   123,
			"name": "John Doe",
			"type": "cached_user",
		}),
	}

	apiV1.FastGET("/cached-user/:id", func(c *gin.Context) {
		userID := c.Param("id")
		if cached, ok := userCache[userID]; ok {
			c.PreMarshaledJSON(200, cached)
		} else {
			c.FastError()
		}
	})

	fmt.Println("✅ All optimization features configured!")
	fmt.Println("\n🧪 Test Endpoints:")
	fmt.Println("   GET /health              - Health check")
	fmt.Println("   GET /ping                - Ping endpoint")
	fmt.Println("   GET /version             - Version info")
	fmt.Println("   GET /api/v1/fast-ok      - Fast OK response")
	fmt.Println("   GET /api/v1/fast-success - Fast success response")
	fmt.Println("   GET /api/v1/fast-pong    - Fast pong response")
	fmt.Println("   GET /api/v1/fast-error   - Fast error response")
	fmt.Println("   GET /api/v1/fast-number  - Fast JSON number")
	fmt.Println("   GET /api/v1/fast-json    - Fast JSON response")
	fmt.Println("   GET /api/v1/string-fast  - Fast string response")
	fmt.Println("   GET /api/v1/pre-marshaled - Pre-marshaled JSON")
	fmt.Println("   GET /api/v1/cached-user/123 - Cached user response")
	fmt.Println("   GET /stats               - Runtime statistics")
	fmt.Println("   GET /debug/pprof/        - Profiling info")

	// 12. Performance stats monitoring
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()

		for range ticker.C {
			stats := gin.GetRuntimeStats()
			log.Printf("📈 Runtime Stats: Goroutines=%d, MemAlloc=%d bytes, GC=%d cycles",
				stats.Goroutines, stats.MemAlloc, stats.GCCycles)
		}
	}()

	// 13. Start server
	fmt.Println("\n🌐 Starting server on :8080...")
	fmt.Println("   Press Ctrl+C to stop")

	// Optimized server configuration
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    60 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	log.Fatal(server.ListenAndServe())
}
