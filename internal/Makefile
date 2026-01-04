.PHONY: test test-race bench bench-short bench-cpu clean

# ====== TESTS ======

test:
	go test -v -count=1 ./...

test-race:
	go test -race -v -count=1 ./...

# ====== BENCHMARKS ======

bench:
	go test  -bench=. -benchmem -benchtime=1s -count=1 ./...

bench-short:
	go test -bench=. -benchmem -benchtime=100ms -count=1 ./...

bench-cpu:
	GOMAXPROCS=1 go test -bench=. -benchmem -count=1 ./...
	GOMAXPROCS=2 go test -bench=. -benchmem -count=1 ./...
	GOMAXPROCS=4 go test -bench=. -benchmem -count=1 ./...
	GOMAXPROCS=8 go test -bench=. -benchmem -count=1 ./...

# ====== CLEAN ======

clean:
	go clean -testcache
