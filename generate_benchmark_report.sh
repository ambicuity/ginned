#!/bin/bash
# Generate Comprehensive Benchmark Report
# This script creates a detailed benchmark report for the Gin framework

set -e

echo "Generating comprehensive Gin framework benchmark report..."
echo "========================================================="

# Ensure benchmark_update.sh is executable
chmod +x benchmark_update.sh

# Create comprehensive report
echo "# 🚀 Gin Framework Benchmark Report" > benchmark_report.txt
echo "" >> benchmark_report.txt
./benchmark_update.sh >> benchmark_report.txt
echo "" >> benchmark_report.txt

echo "## 📊 Detailed Benchmark Results" >> benchmark_report.txt
echo "" >> benchmark_report.txt
echo "### Core Performance Benchmarks" >> benchmark_report.txt
echo "\`\`\`" >> benchmark_report.txt
go test -bench="BenchmarkOneRoute|BenchmarkRecoveryMiddleware|BenchmarkLoggerMiddleware|BenchmarkOneRouteJSON|BenchmarkGinJSONRenderOptimized" -benchmem -benchtime=2s 2>/dev/null | grep "^Benchmark" >> benchmark_report.txt
echo "\`\`\`" >> benchmark_report.txt
echo "" >> benchmark_report.txt

echo "### All Benchmark Results" >> benchmark_report.txt
echo "\`\`\`" >> benchmark_report.txt
go test -bench=. -benchmem -benchtime=1s 2>/dev/null | grep "^Benchmark" >> benchmark_report.txt
echo "\`\`\`" >> benchmark_report.txt
echo "" >> benchmark_report.txt

echo "### Performance Summary" >> benchmark_report.txt
echo "| Benchmark | Performance | Memory | Allocations |" >> benchmark_report.txt
echo "|-----------|-------------|---------|-------------|" >> benchmark_report.txt
go test -bench="BenchmarkOneRoute|BenchmarkOneRouteJSON|BenchmarkGinJSONRenderOptimized" -benchmem -benchtime=1s 2>/dev/null | grep "^Benchmark" | awk '{print "| " $1 " | " $3 " | " $5 " | " $7 " |"}' >> benchmark_report.txt

echo "✅ Benchmark report generated: benchmark_report.txt"
echo ""
echo "📋 Report preview:"
echo "==================="
head -20 benchmark_report.txt
echo ""
echo "📄 Full report saved to: benchmark_report.txt"