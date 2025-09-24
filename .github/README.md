# GitHub Actions for Gin Framework

This directory contains GitHub Actions workflows for automated testing and benchmarking of the Gin web framework.

## Workflows

### 1. CI (`ci.yml`)
**Purpose**: Basic continuous integration testing
**Triggers**: Push and pull requests to main/master branches
**What it does**:
- Tests on Go versions 1.22, 1.23, and 1.24
- Runs standard tests
- Runs basic benchmarks (500ms each)
- Tests the benchmark automation script

### 2. Benchmark Tests (`benchmarks.yml`)
**Purpose**: Comprehensive benchmark testing and performance monitoring
**Triggers**: 
- Push/PR to main/master when Go files or benchmark files change
- Manual dispatch
**What it does**:
- Tests on Go versions 1.21, 1.22, 1.23, and 1.24
- Runs comprehensive benchmark suite
- Tests the `benchmark_update.sh` script
- Validates benchmark output format
- Performs performance regression checks
- **Generates detailed benchmark reports for all PRs and pushes**
- Comments on PRs with benchmark results summary
- Uploads full benchmark reports as artifacts

### 3. Dependabot (`../dependabot.yml`)
**Purpose**: Automated dependency updates
**What it does**:
- Weekly checks for Go module updates
- Weekly checks for GitHub Actions updates
- Automatically creates PRs for updates

## Benchmark Reporting

The benchmark workflow now generates comprehensive reports that include:

### 📊 Report Sections
- **System Information**: Current environment details
- **Core Performance Benchmarks**: Key performance metrics
- **All Benchmark Results**: Complete benchmark suite
- **Performance Summary**: Tabular comparison of key benchmarks

### 🚀 Automatic PR Comments
For pull requests, the workflow automatically:
- Generates a comprehensive benchmark report
- Posts a summary comment with key performance metrics
- Provides links to full reports in workflow artifacts

### 📋 Manual Report Generation
Use the local script for development:
```bash
./generate_benchmark_report.sh
```

## Benchmark Script Integration

The workflows integrate with the `benchmark_update.sh` script to:
- Validate the script works correctly in CI
- Generate formatted benchmark headers
- Ensure consistency in benchmark reporting
- Provide performance regression detection

## Performance Monitoring

The benchmark workflow includes performance regression checks:
- Monitors `BenchmarkOneRoute` performance
- Alerts if performance degrades significantly (>200ns/op)
- Provides detailed performance metrics in artifacts

## Artifacts

Benchmark reports are uploaded as artifacts with 30-day retention for:
- Historical performance tracking
- Detailed analysis of benchmark results
- Debugging performance issues

## Usage

### Running Benchmarks Manually
```bash
# Run all benchmarks
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkOneRoute -benchmem

# Use the automation script
./benchmark_update.sh
```

### Triggering Workflows
- **Automatic**: Workflows trigger on pushes and PRs
- **Manual**: Use "Actions" tab → "Benchmark Tests" → "Run workflow"

### Viewing Results
- Check "Actions" tab for workflow runs
- Download benchmark reports from artifacts
- View PR comments for benchmark summaries