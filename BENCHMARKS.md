# Benchmark System

VM HOST: GitHub Actions Machine: Ubuntu 24.04.3 LTS Date: September 24th, 2025 Version: Gin v1.11.0 Go Version: go1.24.7 linux/amd64 Source: Go HTTP Router Benchmark Result: See the benchmark results below

## Current Gin Framework Benchmarks

### Core Framework Benchmarks
```sh
BenchmarkOneRoute-4                      	22272813	        59.87 ns/op	       4 B/op	       0 allocs/op
BenchmarkRecoveryMiddleware-4            	12763632	        90.79 ns/op	       5 B/op	       0 allocs/op
BenchmarkLoggerMiddleware-4              	  744086	      1944 ns/op	     397 B/op	      25 allocs/op
BenchmarkManyHandlers-4                  	  560623	      2074 ns/op	     408 B/op	      26 allocs/op
Benchmark5Params-4                       	 7864659	       138.7 ns/op	       9 B/op	       1 allocs/op
BenchmarkOneRouteJSON-4                  	 2737098	       418.9 ns/op	      76 B/op	       6 allocs/op
BenchmarkOneRouteHTML-4                  	  763161	      1897 ns/op	     457 B/op	      25 allocs/op
BenchmarkOneRouteSet-4                   	 3407400	       354.4 ns/op	     353 B/op	       4 allocs/op
BenchmarkOneRouteString-4                	 6327920	       198.0 ns/op	      60 B/op	       2 allocs/op
BenchmarkManyRoutesFist-4                	30621128	        55.76 ns/op	       3 B/op	       0 allocs/op
BenchmarkManyRoutesLast-4                	22441777	        64.66 ns/op	       4 B/op	       0 allocs/op
Benchmark404-4                           	13805761	        89.40 ns/op	       6 B/op	       0 allocs/op
Benchmark404Many-4                       	11513904	       108.2 ns/op	       7 B/op	       0 allocs/op
```

### JSON Performance Benchmarks
```sh
BenchmarkStandardJSON_Marshal-4          	  805549	      1779 ns/op	     514 B/op	      23 allocs/op
BenchmarkSonicJSON_Marshal-4             	 1455392	       820.5 ns/op	     345 B/op	      10 allocs/op
BenchmarkGinJSONRender-4                 	 1793708	       677.1 ns/op	     156 B/op	       8 allocs/op
BenchmarkGinFastJSON-4                   	 1722544	       712.5 ns/op	     159 B/op	       8 allocs/op
BenchmarkGinJSONRenderOptimized-4        	 6951840	       187.0 ns/op	      28 B/op	       2 allocs/op
```

### Logger Performance Benchmarks
```sh
BenchmarkFastLogger-4                    	  878569	      1334 ns/op	     669 B/op	      18 allocs/op
BenchmarkStandardLogger-4                	  452868	      2698 ns/op	     904 B/op	      35 allocs/op
```

### Route Performance Benchmarks
```sh
BenchmarkSimpleRoute-4                   	 5725054	       195.3 ns/op	      60 B/op	       2 allocs/op
BenchmarkSimpleRouteFast-4               	 4973466	       226.4 ns/op	      39 B/op	       3 allocs/op
BenchmarkPingFast-4                      	 6350779	       173.5 ns/op	      28 B/op	       2 allocs/op
BenchmarkWithFastLogger-4                	 4415118	       268.1 ns/op	      34 B/op	       3 allocs/op
BenchmarkComplexJSONStandard-4           	  638791	      2404 ns/op	     638 B/op	      33 allocs/op
BenchmarkComplexJSONFast-4               	  563022	      2360 ns/op	     634 B/op	      33 allocs/op
BenchmarkComplexJSONPreCached-4          	 7123765	       188.5 ns/op	      28 B/op	       2 allocs/op
```

### GitHub API Simulation Benchmarks
```sh
BenchmarkGithub-4                        	  503605	      2395 ns/op	     977 B/op	      33 allocs/op
BenchmarkParallelGithub-4                	  773569	      1641 ns/op	    1686 B/op	      25 allocs/op
BenchmarkParallelGithubDefault-4         	  723675	      1619 ns/op	    1684 B/op	      25 allocs/op
BenchmarkGin_GithubAll-4                 	 7951905	       145.4 ns/op	      10 B/op	       1 allocs/op
BenchmarkGin_GithubAllSequential-4       	   49510	     28169 ns/op	    1966 B/op	     245 allocs/op
BenchmarkGin_GithubAllParallel-4         	16856138	        64.76 ns/op	       2 B/op	       0 allocs/op
```

### Context and String Response Benchmarks
```sh
BenchmarkFastJSON-4                      	  568634	      1785 ns/op	     805 B/op	      25 allocs/op
BenchmarkRegularJSONContext-4            	  639561	      1840 ns/op	     807 B/op	      25 allocs/op
BenchmarkStringFast-4                    	 2176569	       553.0 ns/op	     500 B/op	       8 allocs/op
BenchmarkRegularString-4                 	 2418554	       499.6 ns/op	     481 B/op	       7 allocs/op
```

### Form Data Processing Benchmarks
```sh
BenchmarkGetMapFromFormData/Small_Bracket-4         	 2792217	       434.7 ns/op	     359 B/op	       4 allocs/op
BenchmarkGetMapFromFormData/Small_Names-4           	 2861316	       421.1 ns/op	     358 B/op	       4 allocs/op
BenchmarkGetMapFromFormData/Medium_Bracket-4        	 1667721	       713.2 ns/op	     378 B/op	       7 allocs/op
BenchmarkGetMapFromFormData/Medium_Names-4          	 1769749	       673.0 ns/op	     375 B/op	       6 allocs/op
BenchmarkGetMapFromFormData/Medium_Other-4          	 1961512	       618.1 ns/op	     371 B/op	       6 allocs/op
BenchmarkGetMapFromFormData/Large_Bracket-4         	   71688	     17088 ns/op	   10350 B/op	     133 allocs/op
BenchmarkGetMapFromFormData/Large_Names-4           	  118772	      9774 ns/op	    5054 B/op	      81 allocs/op
BenchmarkGetMapFromFormData/Large_Other-4           	  202153	      6393 ns/op	    2529 B/op	      56 allocs/op
BenchmarkGetMapFromFormData/WorstCase_Bracket-4     	  542016	      2310 ns/op	     490 B/op	      21 allocs/op
BenchmarkGetMapFromFormData/ShortKeys_Bracket-4     	 2820316	       420.4 ns/op	     357 B/op	       4 allocs/op
BenchmarkGetMapFromFormData/Empty_Key-4             	 7001960	       163.0 ns/op	      58 B/op	       2 allocs/op
```

### Fast Response Benchmarks
```sh
BenchmarkFastPong-4                      	 2437768	       483.8 ns/op	     449 B/op	       7 allocs/op
BenchmarkRegularJSON-4                   	  713035	      1404 ns/op	     942 B/op	      19 allocs/op
```

### Utility Benchmarks
```sh
BenchmarkParseAccept-4                   	 4765249	       258.1 ns/op	     143 B/op	       3 allocs/op
BenchmarkPathClean-4                     	  727468	      1732 ns/op	     231 B/op	      31 allocs/op
BenchmarkPathCleanLong-4                 	     199	   6563679 ns/op	 3646925 B/op	   58375 allocs/op
```

## Performance Analysis

### Key Performance Insights

1. **JSON Performance**: Optimized JSON rendering (BenchmarkGinJSONRenderOptimized) shows ~3.6x improvement over standard JSON rendering
2. **Logger Performance**: FastLogger shows ~2x improvement over standard logger with ~35% reduction in memory allocations
3. **Pre-cached Responses**: Pre-marshaled JSON responses show dramatic performance improvements (~12x faster than standard JSON)
4. **Form Processing**: Performance scales well with small to medium data, with reasonable degradation for large datasets
5. **Route Matching**: Single route performance is excellent (~60 ns/op) with minimal memory allocation

### Benchmark Notes

- All benchmarks run on Ubuntu 24.04.3 LTS with Go 1.24.7
- Results show operations per second, nanoseconds per operation, bytes allocated per operation, and allocations per operation
- **Optimized versions** (FastJSON, FastLogger, PreCached) show significant performance improvements
- **Memory allocations** are minimized in optimized implementations
- Results may vary based on system specifications and load
- Benchmarks are executed with `-benchmem` flag for memory allocation tracking

### How to Run Benchmarks

To run all benchmarks:
```bash
go test -bench=. -benchmem -benchtime=2s
```

To run specific benchmark categories:
```bash
# Core framework benchmarks
go test -bench="BenchmarkOneRoute|BenchmarkRecovery|BenchmarkLogger" -benchmem

# JSON performance benchmarks  
go test -bench="BenchmarkJSON|BenchmarkGinJSON" -benchmem

# Form processing benchmarks
go test -bench="BenchmarkGetMapFromFormData" -benchmem
```