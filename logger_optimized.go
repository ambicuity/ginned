// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"io"
	"strconv"
	"sync"
	"time"
)

// FastLogger is a highly optimized logger that reduces allocations
func FastLogger() HandlerFunc {
	return FastLoggerWithWriter(DefaultWriter)
}

// FastLoggerWithWriter creates an optimized logger with specified writer
func FastLoggerWithWriter(out io.Writer) HandlerFunc {
	// Pre-allocate buffer pool for log formatting
	bufferPool := &sync.Pool{
		New: func() interface{} {
			return make([]byte, 0, 256)
		},
	}

	return func(c *Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Skip logging for certain paths if needed
		if path == "/ping" || path == "/health" {
			return
		}

		// Get buffer from pool
		buf := bufferPool.Get().([]byte)
		defer func() {
			bufferPool.Put(buf[:0])
		}()

		// Fast log formatting without fmt.Sprintf
		latency := time.Since(start)
		
		buf = append(buf, "[GIN] "...)
		buf = append(buf, start.Format("2006/01/02 - 15:04:05")...)
		buf = append(buf, " | "...)
		buf = strconv.AppendInt(buf, int64(c.Writer.Status()), 10)
		buf = append(buf, " | "...)
		buf = append(buf, latency.String()...)
		buf = append(buf, " | "...)
		buf = append(buf, c.ClientIP()...)
		buf = append(buf, " | "...)
		buf = append(buf, c.Request.Method...)
		buf = append(buf, " "...)
		buf = append(buf, path...)
		
		if raw != "" {
			buf = append(buf, "?"...)
			buf = append(buf, raw...)
		}
		
		if len(c.Errors) > 0 {
			buf = append(buf, " | "...)
			buf = append(buf, c.Errors.String()...)
		}
		
		buf = append(buf, '\n')
		
		out.Write(buf)
	}
}