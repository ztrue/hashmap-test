# Test Work
https://github.com/trafficstars/test-job

`make install` – install dependencies (testify)

`make test` – run unit tests

`make bench` – run benchmarks

```
BenchmarkHashMap/16-SET-8           1000000    1894 ns/op    16 B/op    1 allocs/op
BenchmarkHashMap/16-GET-8           1000000    1828 ns/op     3 B/op    0 allocs/op
BenchmarkHashMap/16-UNSET-8         1000000      99.5 ns/op   3 B/op    0 allocs/op

BenchmarkHashMap/64-SET-8           1000000     802 ns/op    15 B/op    1 allocs/op
BenchmarkHashMap/64-GET-8           1000000     747 ns/op     3 B/op    0 allocs/op
BenchmarkHashMap/64-UNSET-8         1000000      98.5 ns/op   3 B/op    0 allocs/op

BenchmarkHashMap/128-SET-8          1000000     486 ns/op    15 B/op    1 allocs/op
BenchmarkHashMap/128-GET-8          1000000     451 ns/op     3 B/op    0 allocs/op
BenchmarkHashMap/128-UNSET-8        1000000      99.2 ns/op   3 B/op    0 allocs/op

BenchmarkHashMap/1024-SET-8         1000000     214 ns/op    15 B/op    1 allocs/op
BenchmarkHashMap/1024-GET-8         1000000     194 ns/op     3 B/op    0 allocs/op
BenchmarkHashMap/1024-UNSET-8       1000000      99.2 ns/op   3 B/op    0 allocs/op

BenchmarkHashMap/16384-SET-8        1000000     160 ns/op    15 B/op    1 allocs/op
BenchmarkHashMap/16384-GET-8        1000000     142 ns/op     3 B/op    0 allocs/op
BenchmarkHashMap/16384-UNSET-8      1000000     100 ns/op     3 B/op    0 allocs/op

BenchmarkNativeMap/SET-8            1000000     136 ns/op    15 B/op    1 allocs/op
BenchmarkNativeMap/GET-8            1000000     116 ns/op     3 B/op    0 allocs/op
BenchmarkNativeMap/UNSET-8          1000000      69.6 ns/op   3 B/op    0 allocs/op

BenchmarkDefaultHashFunc/16-8       1000000      91.4 ns/op   3 B/op    0 allocs/op
BenchmarkDefaultHashFunc/64-8       1000000      95.9 ns/op   3 B/op    0 allocs/op
BenchmarkDefaultHashFunc/128-8      1000000      90.0 ns/op   3 B/op    0 allocs/op
BenchmarkDefaultHashFunc/1024-8     1000000      91.0 ns/op   3 B/op    0 allocs/op
BenchmarkDefaultHashFunc/16384-8    1000000      91.0 ns/op   3 B/op    0 allocs/op
```
