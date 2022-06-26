ALL: combined-with-oceans.json bin/render-one bin/render-many

bin/%: cmd/%/main.go $(shell find pkg -type f)
	CGO_CFLAGS="-I$(shell brew --prefix cairo)/include" go build -o $@ ./cmd/$*

combined-with-oceans.json:
	curl -OL https://github.com/evansiroky/timezone-boundary-builder/releases/download/2021c/timezones-with-oceans.geojson.zip
	unzip timezones-with-oceans.geojson.zip
	rm -f timezones-with-oceans.geojson.zip

clean:
	rm -rf bin

real-clean: clean
	rm -f combined-with-oceans.json
