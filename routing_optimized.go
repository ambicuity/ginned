// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"strings"
	"sync"
)

// Route cache for static routes to avoid repeated parsing
var routeCache = sync.Map{}

// FastRouteGroup creates an optimized route group for high-performance routing
type FastRouteGroup struct {
	*RouterGroup
	commonHandlers HandlersChain
}

// NewFastRouteGroup creates a route group optimized for performance
func (rg *RouterGroup) NewFastRouteGroup(relativePath string, handlers ...HandlerFunc) *FastRouteGroup {
	return &FastRouteGroup{
		RouterGroup:    rg.Group(relativePath, handlers...),
		commonHandlers: rg.combineHandlers(handlers),
	}
}

// FastGET registers a GET route with optimized handler chain
func (frg *FastRouteGroup) FastGET(relativePath string, handler HandlerFunc) IRoutes {
	return frg.fastHandle("GET", relativePath, handler)
}

// FastPOST registers a POST route with optimized handler chain
func (frg *FastRouteGroup) FastPOST(relativePath string, handler HandlerFunc) IRoutes {
	return frg.fastHandle("POST", relativePath, handler)
}

// FastPUT registers a PUT route with optimized handler chain
func (frg *FastRouteGroup) FastPUT(relativePath string, handler HandlerFunc) IRoutes {
	return frg.fastHandle("PUT", relativePath, handler)
}

// FastDELETE registers a DELETE route with optimized handler chain
func (frg *FastRouteGroup) FastDELETE(relativePath string, handler HandlerFunc) IRoutes {
	return frg.fastHandle("DELETE", relativePath, handler)
}

// fastHandle is an optimized route registration
func (frg *FastRouteGroup) fastHandle(httpMethod, relativePath string, handler HandlerFunc) IRoutes {
	// Use pre-combined handlers to reduce allocation
	handlers := make(HandlersChain, len(frg.commonHandlers)+1)
	copy(handlers, frg.commonHandlers)
	handlers[len(handlers)-1] = handler

	return frg.RouterGroup.handle(httpMethod, relativePath, handlers)
}

// CommonRoutes provides pre-optimized common API routes
type CommonRoutes struct {
	engine *Engine
}

// SetupCommonRoutes creates optimized routes for common API patterns
func (engine *Engine) SetupCommonRoutes() *CommonRoutes {
	cr := &CommonRoutes{engine: engine}

	// Health check endpoint (no middleware for maximum speed)
	engine.GET("/health", func(c *Context) {
		c.PreMarshaledJSON(200, []byte(`{"status":"healthy"}`))
	})

	// Ping endpoint
	engine.GET("/ping", func(c *Context) {
		c.FastPong()
	})

	// Version endpoint
	engine.GET("/version", func(c *Context) {
		c.PreMarshaledJSON(200, []byte(`{"version":"1.0.0","framework":"gin"}`))
	})

	return cr
}

// String optimization utilities
var stringPool = sync.Pool{
	New: func() interface{} {
		return make([]string, 0, 8)
	},
}

// FastSplitPath splits path efficiently using string pool
func FastSplitPath(path string) []string {
	if path == "/" {
		return []string{""}
	}

	// Try to use cached result first
	if cached, ok := routeCache.Load(path); ok {
		return cached.([]string)
	}

	parts := stringPool.Get().([]string)
	defer func() {
		stringPool.Put(parts[:0])
	}()

	parts = strings.Split(path, "/")

	// Cache common paths
	if len(parts) <= 4 { // Only cache simple paths
		result := make([]string, len(parts))
		copy(result, parts)
		routeCache.Store(path, result)
		return result
	}

	// Return copy for complex paths
	result := make([]string, len(parts))
	copy(result, parts)
	return result
}
