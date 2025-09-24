// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastPong(t *testing.T) {
	router := New()
	router.GET("/test", func(c *Context) {
		c.FastPong()
	})

	w := PerformRequest(router, http.MethodGet, "/test")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func TestFastOk(t *testing.T) {
	router := New()
	router.GET("/test", func(c *Context) {
		c.FastOk()
	})

	w := PerformRequest(router, http.MethodGet, "/test")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"status":"ok"}`, w.Body.String())
}

func TestFastSuccess(t *testing.T) {
	router := New()
	router.GET("/test", func(c *Context) {
		c.FastSuccess()
	})

	w := PerformRequest(router, http.MethodGet, "/test")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"status":"success"}`, w.Body.String())
}

func TestFastError(t *testing.T) {
	router := New()
	router.GET("/test", func(c *Context) {
		c.FastError()
	})

	w := PerformRequest(router, http.MethodGet, "/test")

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"status":"error"}`, w.Body.String())
}

func TestFastJSONNumber(t *testing.T) {
	router := New()
	router.GET("/test", func(c *Context) {
		c.FastJSONNumber(http.StatusOK, 42)
	})

	w := PerformRequest(router, http.MethodGet, "/test")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"value":42}`, w.Body.String())
}

func TestGetOrSetCommonJSON(t *testing.T) {
	// Test caching functionality
	obj := map[string]string{"test": "value"}

	// First call should marshal and cache
	result1 := GetOrSetCommonJSON("test-key", obj)
	assert.NotNil(t, result1)
	assert.Contains(t, string(result1), "test")
	assert.Contains(t, string(result1), "value")

	// Second call should return cached result
	result2 := GetOrSetCommonJSON("test-key", obj)
	assert.Equal(t, result1, result2)
}

func TestPreMarshaledJSON(t *testing.T) {
	jsonData := []byte(`{"custom":"response"}`)

	router := New()
	router.GET("/test", func(c *Context) {
		c.PreMarshaledJSON(http.StatusCreated, jsonData)
	})

	w := PerformRequest(router, http.MethodGet, "/test")

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, `{"custom":"response"}`, w.Body.String())
}

// Benchmark tests to verify performance
func BenchmarkFastPong(b *testing.B) {
	router := New()
	router.GET("/ping", func(c *Context) {
		c.FastPong()
	})

	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w := newMockWriter()
		router.ServeHTTP(w, req)
	}
}

func BenchmarkRegularJSON(b *testing.B) {
	router := New()
	router.GET("/ping", func(c *Context) {
		c.JSON(http.StatusOK, H{"message": "pong"})
	})

	req, _ := http.NewRequest(http.MethodGet, "/ping", nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w := newMockWriter()
		router.ServeHTTP(w, req)
	}
}
