// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"encoding/json"
	"testing"

	"github.com/bytedance/sonic"
)

type TestData struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
	Data   interface{} `json:"data"`
}

var testDataObj = TestData{
	Status: "success",
	Code:   200,
	Data: map[string]interface{}{
		"message": "Hello World",
		"user":    "test_user",
		"items":   []string{"item1", "item2", "item3"},
	},
}

// Benchmark standard encoding/json
func BenchmarkStandardJSON_Marshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(testDataObj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark sonic JSON
func BenchmarkSonicJSON_Marshal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := sonic.Marshal(testDataObj)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// Benchmark standard Gin JSON rendering
func BenchmarkGinJSONRender(b *testing.B) {
	router := New()
	data := TestData{Status: "ok", Code: 200, Data: "test"}
	router.GET("/json", func(c *Context) {
		c.JSON(200, data)
	})
	runRequest(b, router, "GET", "/json")
}

// Benchmark Fast JSON rendering (our optimization)
func BenchmarkGinFastJSON(b *testing.B) {
	router := New()
	data := TestData{Status: "ok", Code: 200, Data: "test"}
	router.GET("/json", func(c *Context) {
		c.FastJSON(200, data)
	})
	runRequest(b, router, "GET", "/json")
}

// Benchmark optimized Gin JSON rendering (pre-marshaled)
func BenchmarkGinJSONRenderOptimized(b *testing.B) {
	router := New()
	// Pre-marshal JSON once for static responses
	jsonBytes := []byte(`{"status":"ok","code":200,"data":"test"}`)
	router.GET("/json", func(c *Context) {
		c.PreMarshaledJSON(200, jsonBytes)
	})
	runRequest(b, router, "GET", "/json")
}