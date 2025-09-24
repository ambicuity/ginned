// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"strconv"
	"sync"

	"github.com/gin-gonic/gin/codec/json"
)

// Common JSON responses that can be pre-marshaled
var (
	commonJSONResponses = sync.Map{}

	// Pre-marshaled common responses
	JSONSuccess = []byte(`{"status":"success"}`)
	JSONError   = []byte(`{"status":"error"}`)
	JSONPong    = []byte(`{"message":"pong"}`)
	JSONOk      = []byte(`{"status":"ok"}`)
)

// GetOrSetCommonJSON returns cached JSON or creates and caches it
func GetOrSetCommonJSON(key string, obj interface{}) []byte {
	if cached, ok := commonJSONResponses.Load(key); ok {
		return cached.([]byte)
	}

	// Marshal and cache the response
	marshaled, err := json.API.Marshal(obj)
	if err != nil {
		return nil
	}

	commonJSONResponses.Store(key, marshaled)
	return marshaled
}

// FastPong renders the common ping/pong response with zero marshaling
func (c *Context) FastPong() {
	c.PreMarshaledJSON(200, JSONPong)
}

// FastOk renders a simple OK response
func (c *Context) FastOk() {
	c.PreMarshaledJSON(200, JSONOk)
}

// FastSuccess renders a success response
func (c *Context) FastSuccess() {
	c.PreMarshaledJSON(200, JSONSuccess)
}

// FastError renders an error response
func (c *Context) FastError() {
	c.PreMarshaledJSON(500, JSONError)
}

// Pool for byte buffers to reduce allocations in string formatting
var bufferPool = sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 64)
	},
}

// Fast integer to JSON response
func (c *Context) FastJSONNumber(code int, number int64) {
	buf := bufferPool.Get().([]byte)
	defer bufferPool.Put(buf[:0])

	buf = append(buf, `{"value":`...)
	buf = strconv.AppendInt(buf, number, 10)
	buf = append(buf, '}')

	c.PreMarshaledJSON(code, buf)
}
