// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFastJSON(t *testing.T) {
	testObj := map[string]interface{}{
		"name": "test",
		"age":  25,
	}

	router := New()
	router.GET("/test", func(c *Context) {
		c.FastJSON(http.StatusOK, testObj)
	})

	w := PerformRequest(router, http.MethodGet, "/test")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Contains(t, w.Body.String(), "test")
	assert.Contains(t, w.Body.String(), "25")
}

func TestStringFast(t *testing.T) {
	testText := "Hello, World!"

	router := New()
	router.GET("/test", func(c *Context) {
		c.StringFast(http.StatusOK, testText)
	})

	w := PerformRequest(router, http.MethodGet, "/test")

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
	assert.Equal(t, testText, w.Body.String())
}

func TestFastJSONPanic(t *testing.T) {
	// Test that FastJSON panics on marshal error
	// Create an object that cannot be marshaled to JSON
	badObj := make(chan int) // channels cannot be marshaled to JSON

	assert.Panics(t, func() {
		router := New()
		router.GET("/test", func(c *Context) {
			c.FastJSON(http.StatusOK, badObj)
		})
		w := PerformRequest(router, http.MethodGet, "/test")
		_ = w
	})
}

// Benchmark tests to compare performance
func BenchmarkFastJSON(b *testing.B) {
	router := New()
	testObj := map[string]interface{}{
		"message": "hello",
		"status":  "ok",
		"count":   42,
	}

	router.GET("/test", func(c *Context) {
		c.FastJSON(http.StatusOK, testObj)
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w := newMockWriter()
		router.ServeHTTP(w, req)
	}
}

func BenchmarkRegularJSONContext(b *testing.B) {
	router := New()
	testObj := map[string]interface{}{
		"message": "hello",
		"status":  "ok",
		"count":   42,
	}

	router.GET("/test", func(c *Context) {
		c.JSON(http.StatusOK, testObj)
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w := newMockWriter()
		router.ServeHTTP(w, req)
	}
}

func BenchmarkStringFast(b *testing.B) {
	router := New()
	testText := "Hello, World! This is a test string."

	router.GET("/test", func(c *Context) {
		c.StringFast(http.StatusOK, testText)
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w := newMockWriter()
		router.ServeHTTP(w, req)
	}
}

func BenchmarkRegularString(b *testing.B) {
	router := New()
	testText := "Hello, World! This is a test string."

	router.GET("/test", func(c *Context) {
		c.String(http.StatusOK, testText)
	})

	req, _ := http.NewRequest(http.MethodGet, "/test", nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		w := newMockWriter()
		router.ServeHTTP(w, req)
	}
}