## Intro
Simple non-recursive realisation of a splay tree.

## Import
 "github.com/awskii/splay"

## Key
By default, Key has type `int`, but you can change to fit your needs.
Key should be comparable.

## Tests
```
$ go test -cover

PASS
coverage: 94.6% of statements
ok  	github.com/awskii/splay	0.180s
```

## Benchmarks
```
$ go test -bench=. -benchmem

goos: darwin
goarch: amd64
pkg: github.com/awskii/splay

(4 CPU)                    (Elements)       (time/op)             (memory overhead)
BenchmarkInsertIdle-4     	20000000	      67.8 ns/op	      32 B/op	       1 allocs/op
BenchmarkSearchIdle-4     	50000000	      21.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkInsertRandom-4   	 1000000	      1576 ns/op	      32 B/op	       1 allocs/op
BenchmarkSearchRandom-4   	 1000000	      1525 ns/op	       0 B/op	       0 allocs/op
```
