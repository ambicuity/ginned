// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastLogger(t *testing.T) {
	buffer := new(bytes.Buffer)
	router := New()
	router.Use(FastLoggerWithWriter(buffer))
	router.GET("/test", func(c *Context) {
		c.String(http.StatusOK, "test")
	})

	w := PerformRequest(router, http.MethodGet, "/test")
	assert.Equal(t, http.StatusOK, w.Code)

	// Check that log was written
	logOutput := buffer.String()
	assert.Contains(t, logOutput, "[GIN]")
	assert.Contains(t, logOutput, "200")
	assert.Contains(t, logOutput, "GET")
	assert.Contains(t, logOutput, "/test")
}

func TestFastLoggerSkipsPingAndHealth(t *testing.T) {
	buffer := new(bytes.Buffer)
	router := New()
	router.Use(FastLoggerWithWriter(buffer))

	router.GET("/ping", func(c *Context) {
		c.String(http.StatusOK, "pong")
	})

	router.GET("/health", func(c *Context) {
		c.String(http.StatusOK, "healthy")
	})

	// Test /ping - should not be logged
	w := PerformRequest(router, http.MethodGet, "/ping")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Empty(t, buffer.String())

	// Reset buffer
	buffer.Reset()

	// Test /health - should not be logged
	w = PerformRequest(router, http.MethodGet, "/health")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Empty(t, buffer.String())
}

func TestFastLoggerWithQuery(t *testing.T) {
	buffer := new(bytes.Buffer)
	router := New()
	router.Use(FastLoggerWithWriter(buffer))
	router.GET("/test", func(c *Context) {
		c.String(http.StatusOK, "test")
	})

	w := PerformRequest(router, http.MethodGet, "/test?foo=bar&baz=qux")
	assert.Equal(t, http.StatusOK, w.Code)

	// Check that query parameters are logged
	logOutput := buffer.String()
	assert.Contains(t, logOutput, "?foo=bar&baz=qux")
}

func TestFastLoggerWithErrors(t *testing.T) {
	buffer := new(bytes.Buffer)
	router := New()
	router.Use(FastLoggerWithWriter(buffer))
	router.GET("/error", func(c *Context) {
		c.Error(assert.AnError)
		c.String(http.StatusInternalServerError, "error")
	})

	w := PerformRequest(router, http.MethodGet, "/error")
	assert.Equal(t, http.StatusInternalServerError, w.Code)

	// Check that error is logged
	logOutput := buffer.String()
	assert.Contains(t, logOutput, "500")
	assert.Contains(t, logOutput, "assert.AnError")
}

// Benchmark to show performance improvement
func BenchmarkFastLogger(b *testing.B) {
	buffer := new(bytes.Buffer)
	router := New()
	router.Use(FastLoggerWithWriter(buffer))
	router.GET("/test", func(c *Context) {
		c.String(http.StatusOK, "test")
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w := newMockWriter()
		router.ServeHTTP(w, req)
		buffer.Reset() // Reset buffer to avoid growing infinitely
	}
}

func BenchmarkStandardLogger(b *testing.B) {
	buffer := new(bytes.Buffer)
	router := New()
	router.Use(LoggerWithWriter(buffer))
	router.GET("/test", func(c *Context) {
		c.String(http.StatusOK, "test")
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w := newMockWriter()
		router.ServeHTTP(w, req)
		buffer.Reset() // Reset buffer to avoid growing infinitely
	}
}
