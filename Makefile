GOSRC:=$(shell find . -name '*.go')

.PHONY: cpuprof

# → https://blog.golang.org/cover
cover: coverage.html

cpuprof:
	go test --cpuprofile=cpu.prof --bench=.

memprof:
	go test --memprofile=mem.prof --bench=.

show-cpuprof: cpuprof
	go tool pprof :6060 groph.test cpu.prof

show-memprof: memprof
	go tool pprof :6060 groph.test mem.prof

coverage.html: coverage.out
	go tool cover -html=$< -o $@

coverage.out: $(GOSRC)
	go test -coverprofile=$@ ./... || true
#	go test -covermode=count -coverprofile=$@ || true
