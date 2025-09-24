#!/bin/bash
# Benchmark Update Script for Gin Framework
# This script generates the benchmark header format for BENCHMARKS.md

set -e

echo "Gin Framework Benchmark Update Script"
echo "======================================"

# Get system information
VM_HOST="GitHub Actions"
MACHINE=$(lsb_release -d 2>/dev/null | cut -f2 || echo "Ubuntu $(cat /etc/os-release | grep VERSION= | cut -d'"' -f2)")
DATE=$(date '+%B %dth, %Y')
VERSION=$(grep 'const Version' version.go | cut -d'"' -f2)
GO_VERSION=$(go version | cut -d' ' -f3,4)

echo "System Information:"
echo "VM HOST: $VM_HOST"
echo "Machine: $MACHINE"
echo "Date: $DATE"
echo "Version: Gin $VERSION"
echo "Go Version: $GO_VERSION"
echo ""

# Create benchmark header in the requested format
BENCHMARK_HEADER="VM HOST: $VM_HOST Machine: $MACHINE Date: $DATE Version: Gin $VERSION Go Version: $GO_VERSION Source: Go HTTP Router Benchmark Result: See the benchmark results below"

echo "Benchmark header for BENCHMARKS.md:"
echo "===================================="
echo "$BENCHMARK_HEADER"
echo ""

echo "To run benchmarks manually, use:"
echo "go test -bench=. -benchmem -benchtime=2s 2>/dev/null | grep '^Benchmark'"
echo ""

# Quick test of a single benchmark
echo "Sample benchmark (BenchmarkOneRoute):"
go test -bench=BenchmarkOneRoute -benchtime=1s 2>/dev/null | grep "^BenchmarkOneRoute" || echo "Benchmark test failed"