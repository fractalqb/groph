GOSRC:=$(shell find . -name '*.go')

.PHONY: cpuprof

README.md: README.md~
	cp $< $@

# → https://blog.golang.org/cover
cover: coverage.html

cpuprof:
	go test --cpuprofile=cpu.prof --bench=.

show-cpuprof: cpuprof
	go tool pprof :6060 groph.test cpu.prof

coverage.html: coverage.out
	go tool cover -html=$< -o $@

coverage.out: $(GOSRC)
	go test -coverprofile=$@ ./... || true
#	go test -covermode=count -coverprofile=$@ || true

cov=$(shell go tool cover -func=coverage.out \
            | egrep '^total:' \
            | awk '{print $$3}' \
            | tr "%" " ")

README.md~: coverage.html
	awk -v cov=$(cov) '/^\[!\[Test Coverage]/{ \
        cov=sprintf("%.0f", cov); \
		printf "[![Test Coverage](https://img.shields.io/badge/coverage-"; \
		printf "%s%%25-", cov ;\
		if (cov < 50) { \
			printf "red" \
	    } else if (cov < 75) { \
			printf "orange" \
	    } else if (cov < 90) { \
			printf "yellow" \
	    } else { \
			printf "green" \
		}; \
        print ".svg)](file:coverage.html)" ;\
     	next \
	} \
	{ print }' README.md > $@
