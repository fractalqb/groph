all: plain.png fancy.png

%.png: %.dot
	dot -Tpng $< > $@

%.dot: main.go
	go run main.go
