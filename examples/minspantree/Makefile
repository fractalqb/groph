all: mst-dir.svg mst-undir.svg

%.svg: %.dot
	dot -Tsvg $< > $@

%-dir.dot: minspantree
	./$< > $@

%-undir.dot: minspantree
	./$< -u > $@

minspantree: minspantree.go
	go build
