IMAGES := pos-sm.svg \
	pos-lg.svg \
	pos-sm-summary.svg \
	pos-lg-summary.svg \
	neg-sm.svg \
	neg-lg.svg \
	neg-sm-summary.svg \
	neg-lg-summary.svg


ALL: $(IMAGES) bin/server

.PRECIOUS: bin/render

%.svg: bin/render
	bin/render

bin/%: cmd/%/main.go $(shell find pkg -type f)
	go build -o $@ ./cmd/$*

clean:
	rm -rf bin $(IMAGES)