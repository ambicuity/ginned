// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gin-gonic/gin"
)

// EndpointTester demonstrates all optimization features with test requests
func EndpointTester() {
	fmt.Println("🧪 Testing All Optimized Endpoints")
	fmt.Println("==================================")

	// Setup the same router as the demo
	gin.SetMode(gin.TestMode)
	router := setupOptimizedRouter()

	// Test all endpoints
	testEndpoints := []struct {
		method      string
		path        string
		expected    int
		description string
	}{
		{"GET", "/health", 200, "Health check endpoint"},
		{"GET", "/ping", 200, "Ping endpoint"},
		{"GET", "/version", 200, "Version endpoint"},
		{"GET", "/api/v1/fast-ok", 200, "Fast OK response"},
		{"GET", "/api/v1/fast-success", 200, "Fast success response"},
		{"GET", "/api/v1/fast-pong", 200, "Fast pong response"},
		{"GET", "/api/v1/fast-error", 500, "Fast error response"},
		{"GET", "/api/v1/fast-number", 200, "Fast JSON number"},
		{"GET", "/api/v1/fast-json", 200, "Fast JSON response"},
		{"GET", "/api/v1/string-fast", 200, "Fast string response"},
		{"GET", "/api/v1/pre-marshaled", 200, "Pre-marshaled JSON"},
		{"GET", "/api/v1/cached-user/123", 200, "Cached user response"},
		{"GET", "/stats", 200, "Runtime statistics"},
	}

	fmt.Printf("Testing %d endpoints...\n\n", len(testEndpoints))

	successCount := 0
	for i, test := range testEndpoints {
		success := testEndpoint(router, test.method, test.path, test.expected, test.description)
		if success {
			successCount++
		}
		if i < len(testEndpoints)-1 {
			time.Sleep(10 * time.Millisecond) // Small delay for readability
		}
	}

	fmt.Printf("\n📊 Results: %d/%d endpoints working correctly (%.1f%% success rate)\n",
		successCount, len(testEndpoints), float64(successCount)/float64(len(testEndpoints))*100)

	if successCount == len(testEndpoints) {
		fmt.Println("✅ All optimization features are working perfectly!")
	} else {
		fmt.Printf("⚠️  %d endpoints need attention\n", len(testEndpoints)-successCount)
	}
}

func setupOptimizedRouter() *gin.Engine {
	// Create router with all optimizations (same as demo)
	router := gin.New()

	// Buffer to capture logs for testing
	var logBuffer bytes.Buffer
	router.Use(gin.FastLoggerWithWriter(&logBuffer))
	router.Use(gin.Recovery())

	// Setup common routes
	router.SetupCommonRoutes()
	router.RuntimeStatsEndpoint()

	// Fast route groups
	apiV1 := router.NewFastRouteGroup("/api/v1")

	// Fast response methods
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

	// Optimized context methods
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

	// Pre-marshaled responses
	successResponse := []byte(`{"status":"success","optimization":"pre-marshaled","speed":"maximum"}`)
	apiV1.FastGET("/pre-marshaled", func(c *gin.Context) {
		c.PreMarshaledJSON(200, successResponse)
	})

	// Cached responses
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

	return router
}

func testEndpoint(router *gin.Engine, method, path string, expectedStatus int, description string) bool {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	success := w.Code == expectedStatus
	status := "✅"
	if !success {
		status = "❌"
	}

	fmt.Printf("%s %s %s -> %d (expected %d) - %s\n",
		status, method, path, w.Code, expectedStatus, description)

	// Show response body for interesting endpoints
	if success && (path == "/api/v1/fast-json" || path == "/stats" || path == "/api/v1/cached-user/123") {
		body := w.Body.String()
		if len(body) > 100 {
			body = body[:100] + "..."
		}

		// Try to pretty print JSON
		var jsonData interface{}
		if json.Unmarshal(w.Body.Bytes(), &jsonData) == nil {
			if prettyJSON, err := json.MarshalIndent(jsonData, "  ", "  "); err == nil && len(prettyJSON) < 200 {
				fmt.Printf("  Response: %s\n", string(prettyJSON))
			} else {
				fmt.Printf("  Response: %s\n", body)
			}
		} else {
			fmt.Printf("  Response: %s\n", body)
		}
	}

	return success
}

func main() {
	fmt.Println("🚀 Gin Framework Optimization Demonstration")
	fmt.Println("==========================================")

	// Show build info
	buildInfo := gin.GetBuildInfo()
	fmt.Printf("📊 Build Information:\n")
	fmt.Printf("   Mode: %s\n", buildInfo.Mode)
	fmt.Printf("   JSON Provider: %s\n", buildInfo.JSONProvider)
	fmt.Printf("   Optimized Routes: %t\n", buildInfo.OptimizedRoutes)
	fmt.Printf("   Pooling Enabled: %t\n\n", buildInfo.PoolingEnabled)

	// Test all endpoints
	EndpointTester()
}
