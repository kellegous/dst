ALL: bin/tzimg bin/render-one bin/render-many bin/color-export

bin/%: cmd/%/main.go $(shell find pkg -type f)
	CGO_CFLAGS="-I$(shell brew --prefix cairo)/include" go build -o $@ ./cmd/$*

clean:
	rm -rf bin