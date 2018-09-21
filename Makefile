SOURCES =  \
  main.go  \
  rebed.go \
  rebin/track.go

.PHONY: clean test depends

rebed: $(SOURCES) depends
	go build -o $@

clean:
	rm -f rebed

test: depends
	go test ./...

depends:
	go get ./...
