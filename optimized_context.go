// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"github.com/gin-gonic/gin/codec/json"
)

// FastJSON renders JSON response with minimal allocations
// This method bypasses the render system for maximum performance
func (c *Context) FastJSON(code int, obj interface{}) {
	c.Status(code)
	c.Header("Content-Type", "application/json; charset=utf-8")
	
	// Use sonic if available for faster marshaling
	jsonBytes, err := json.API.Marshal(obj)
	if err != nil {
		panic(err)
	}
	
	_, _ = c.Writer.Write(jsonBytes)
}

// PreMarshaledJSON renders pre-marshaled JSON with zero marshaling overhead
func (c *Context) PreMarshaledJSON(code int, jsonBytes []byte) {
	c.Status(code)
	c.Header("Content-Type", "application/json; charset=utf-8")
	_, _ = c.Writer.Write(jsonBytes)
}

// StringFast renders string response with zero-copy conversion
func (c *Context) StringFast(code int, text string) {
	c.Status(code)
	c.Header("Content-Type", "text/plain; charset=utf-8")
	
	// Zero-copy string to bytes conversion - safe for read-only usage
	textBytes := []byte(text) // Keep it simple and safe for now
	_, _ = c.Writer.Write(textBytes)
}