// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupCommonRoutes(t *testing.T) {
	router := New()
	router.SetupCommonRoutes()

	// Test health endpoint
	w := PerformRequest(router, http.MethodGet, "/health")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"status":"healthy"}`, w.Body.String())

	// Test ping endpoint
	w = PerformRequest(router, http.MethodGet, "/ping")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())

	// Test version endpoint
	w = PerformRequest(router, http.MethodGet, "/version")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"version":"1.0.0","framework":"gin"}`, w.Body.String())
}

func TestNewFastRouteGroup(t *testing.T) {
	router := New()

	// Create a fast route group
	api := router.NewFastRouteGroup("/api/v1")

	// Test FastGET
	api.FastGET("/users", func(c *Context) {
		c.FastOk()
	})

	w := PerformRequest(router, http.MethodGet, "/api/v1/users")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())
}

func TestFastRouteGroupMethods(t *testing.T) {
	router := New()
	api := router.NewFastRouteGroup("/api")

	// Test all HTTP methods
	api.FastGET("/get", func(c *Context) { c.FastSuccess() })
	api.FastPOST("/post", func(c *Context) { c.FastSuccess() })
	api.FastPUT("/put", func(c *Context) { c.FastSuccess() })
	api.FastDELETE("/delete", func(c *Context) { c.FastSuccess() })

	// Test GET
	w := PerformRequest(router, http.MethodGet, "/api/get")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"success"}`, w.Body.String())

	// Test POST
	w = PerformRequest(router, http.MethodPost, "/api/post")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"success"}`, w.Body.String())

	// Test PUT
	w = PerformRequest(router, http.MethodPut, "/api/put")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"success"}`, w.Body.String())

	// Test DELETE
	w = PerformRequest(router, http.MethodDelete, "/api/delete")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"status":"success"}`, w.Body.String())
}

func TestFastSplitPath(t *testing.T) {
	// Test root path
	parts := FastSplitPath("/")
	assert.Equal(t, []string{""}, parts)

	// Test simple path
	parts = FastSplitPath("/api/v1/users")
	expected := []string{"", "api", "v1", "users"}
	assert.Equal(t, expected, parts)

	// Test caching - call twice to ensure cache works
	parts1 := FastSplitPath("/api/v1")
	parts2 := FastSplitPath("/api/v1")
	assert.Equal(t, parts1, parts2)
}

func TestFastRouteGroupWithMiddleware(t *testing.T) {
	router := New()

	// Add middleware to the route group
	api := router.NewFastRouteGroup("/api", func(c *Context) {
		c.Header("X-API-Version", "v1")
		c.Next()
	})

	api.FastGET("/test", func(c *Context) {
		c.FastOk()
	})

	w := PerformRequest(router, http.MethodGet, "/api/test")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "v1", w.Header().Get("X-API-Version"))
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())
}
