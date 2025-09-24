// Copyright 2025 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package gin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBuildInfo(t *testing.T) {
	buildInfo := GetBuildInfo()

	// Verify build info structure
	assert.NotEmpty(t, buildInfo.Mode)
	assert.NotEmpty(t, buildInfo.JSONProvider)
	assert.True(t, buildInfo.OptimizedRoutes)
	assert.True(t, buildInfo.PoolingEnabled)

	// Mode should be set
	assert.Contains(t, []string{"debug", "release", "test"}, buildInfo.Mode)

	// JSON provider should be sonic or another valid provider
	assert.NotEmpty(t, buildInfo.JSONProvider)
}

func TestBuildOptimizationConstants(t *testing.T) {
	// Test that constants are defined
	assert.NotEmpty(t, OptimalGOMAXPROCS)
	assert.NotEmpty(t, OptimalBuildFlags)
	assert.NotEmpty(t, OptimalServerConfig)

	// Test content of constants
	assert.Contains(t, OptimalGOMAXPROCS, "GOMAXPROCS")
	assert.Contains(t, OptimalBuildFlags, "-ldflags")
	assert.Contains(t, OptimalServerConfig, "ReadTimeout")
}
