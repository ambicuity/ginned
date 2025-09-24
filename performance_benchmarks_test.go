// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"testing"
)

// Benchmark simple route without any middleware
func BenchmarkSimpleRoute(b *testing.B) {
	router := New()
	router.GET("/ping", func(c *Context) {
		c.String(200, "pong")
	})
	runRequest(b, router, "GET", "/ping")
}

// Benchmark simple route with fast string response
func BenchmarkSimpleRouteFast(b *testing.B) {
	router := New()
	router.GET("/ping", func(c *Context) {
		c.StringFast(200, "pong")
	})
	runRequest(b, router, "GET", "/ping")
}

// Benchmark ping endpoint with pre-marshaled JSON
func BenchmarkPingFast(b *testing.B) {
	router := New()
	router.GET("/ping", func(c *Context) {
		c.FastPong()
	})
	runRequest(b, router, "GET", "/ping")
}

// Benchmark with standard logger
func BenchmarkWithStandardLogger(b *testing.B) {
	router := New()
	router.Use(Logger())
	router.GET("/ping", func(c *Context) {
		c.FastPong()
	})
	runRequest(b, router, "GET", "/ping")
}

// Benchmark with optimized logger
func BenchmarkWithFastLogger(b *testing.B) {
	router := New()
	router.Use(FastLogger())
	router.GET("/ping", func(c *Context) {
		c.FastPong()
	})
	runRequest(b, router, "GET", "/ping")
}

// Benchmark JSON response with complex data
func BenchmarkComplexJSONStandard(b *testing.B) {
	router := New()
	data := map[string]interface{}{
		"status": "success",
		"data": map[string]interface{}{
			"user":  "test_user",
			"items": []string{"item1", "item2", "item3"},
			"count": 42,
		},
	}
	router.GET("/api/data", func(c *Context) {
		c.JSON(200, data)
	})
	runRequest(b, router, "GET", "/api/data")
}

// Benchmark JSON response with fast JSON
func BenchmarkComplexJSONFast(b *testing.B) {
	router := New()
	data := map[string]interface{}{
		"status": "success",
		"data": map[string]interface{}{
			"user":  "test_user",
			"items": []string{"item1", "item2", "item3"},
			"count": 42,
		},
	}
	router.GET("/api/data", func(c *Context) {
		c.FastJSON(200, data)
	})
	runRequest(b, router, "GET", "/api/data")
}

// Benchmark with pre-cached complex JSON
func BenchmarkComplexJSONPreCached(b *testing.B) {
	router := New()
	// Pre-marshal the JSON response once
	jsonResp := []byte(`{"status":"success","data":{"user":"test_user","items":["item1","item2","item3"],"count":42}}`)
	router.GET("/api/data", func(c *Context) {
		c.PreMarshaledJSON(200, jsonResp)
	})
	runRequest(b, router, "GET", "/api/data")
}

// Benchmark multiple middleware (standard)
func BenchmarkMultipleMiddlewareStandard(b *testing.B) {
	router := New()
	router.Use(Logger(), Recovery())
	router.GET("/api/endpoint", func(c *Context) {
		c.JSON(200, H{"status": "ok"})
	})
	runRequest(b, router, "GET", "/api/endpoint")
}

// Benchmark multiple middleware (optimized)
func BenchmarkMultipleMiddlewareOptimized(b *testing.B) {
	router := New()
	router.Use(FastLogger(), Recovery())
	router.GET("/api/endpoint", func(c *Context) {
		c.FastOk()
	})
	runRequest(b, router, "GET", "/api/endpoint")
}