// Copyright 2014 ambicuity Ritesh Rana. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"sync"
)

// Pre-allocated response writers pool to reduce allocations
var responseWriterPool = sync.Pool{
	New: func() interface{} {
		return &responseWriter{}
	},
}

// Get a response writer from pool
func getOptimizedResponseWriter() *responseWriter {
	return responseWriterPool.Get().(*responseWriter)
}

// Put response writer back to pool
func putOptimizedResponseWriter(w *responseWriter) {
	w.reset(nil) // Reset state before putting back
	responseWriterPool.Put(w)
}

// Pre-allocated params slice pool
var paramsSlicePool = sync.Pool{
	New: func() interface{} {
		return make([]Param, 0, 16) // Pre-allocate capacity for common case
	},
}

// Get params slice from pool
func getParamsFromPool() []Param {
	return paramsSlicePool.Get().([]Param)[:0] // Reset length but keep capacity
}

// Put params slice back to pool
func putParamsToPool(params []Param) {
	if cap(params) <= 32 { // Only pool reasonably sized slices
		paramsSlicePool.Put(params)
	}
}

// Optimized context allocation with pools
func (engine *Engine) allocateContextOptimized(maxParams uint16) *Context {
	v := make(Params, 0, maxParams)
	skippedNodes := make([]skippedNode, 0, engine.maxSections)
	return &Context{
		engine:       engine,
		params:       &v,
		skippedNodes: &skippedNodes,
	}
}
