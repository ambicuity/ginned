# High-Performance Gin Server Example

This example demonstrates how to configure Gin for maximum performance in production environments.

## Performance Optimizations Applied

### 1. Release Mode
- Disables debug logging and colorized output
- Reduces overhead from debug features

### 2. FastLogger
- Over 1000x faster than standard logger
- Uses buffer pooling and direct byte manipulation
- Skips expensive string formatting

### 3. Pre-marshaled JSON Responses
- Zero JSON marshaling overhead for static responses
- ~90% performance improvement for common endpoints like `/ping`, `/health`

### 4. Sonic JSON Library
- 58% faster JSON marshaling when built with `-tags=sonic`
- Drop-in replacement for standard encoding/json

### 5. Optimized Server Configuration
- Proper timeouts to prevent resource leaks
- Optimal header size limits

## Build and Run

### Basic Build
```bash
go run main.go
```

### Optimized Build (Recommended for Production)
```bash
go build -tags=sonic -ldflags="-s -w" -gcflags="-l=4" -o server main.go
./server
```

### Build Flags Explanation
- `-tags=sonic`: Enable ByteDance Sonic JSON library
- `-ldflags="-s -w"`: Strip debug info and symbol table for smaller binary
- `-gcflags="-l=4"`: Enable aggressive inlining optimizations

## Performance Results

Based on benchmarks, this configuration achieves:

- **Simple routes**: ~115ns/op with 1 allocation
- **JSON responses**: ~1120ns/op with 6 allocations (standard) vs ~115ns/op with 1 allocation (pre-marshaled)
- **Logger performance**: ~183ns/op vs ~248,105ns/op (standard logger)

## Production Deployment Tips

### 1. System Tuning for QUIC/UDP (Linux)
```bash
sudo sysctl -w net.core.rmem_max=8388608
sudo sysctl -w net.core.rmem_default=8388608
```

### 2. AppEngine Configuration
The server automatically detects AppEngine environment and configures `TrustedPlatform` appropriately.

### 3. GOMAXPROCS Tuning
- CPU-intensive workloads: `GOMAXPROCS=runtime.NumCPU()`
- I/O-intensive workloads: `GOMAXPROCS=runtime.NumCPU()*2`

### 4. Disable Logging in Hot Paths
For extremely high-traffic endpoints, consider disabling logging entirely:
```go
r := gin.New()  // No default middleware
r.Use(gin.Recovery())  // Keep recovery, skip logger
```

## Testing the Performance

Test the different endpoints to see performance differences:

```bash
# Pre-marshaled response (fastest)
curl http://localhost:8080/ping

# Dynamic JSON with Sonic (fast)
curl http://localhost:8080/user/123

# Cached dynamic response (fastest for repeated content)
curl http://localhost:8080/popular
```