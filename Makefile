SOURCES =  \
  main.go  \
  rebed.go \
  rebin/track.go

.PHONY: clean test

rebed: $(SOURCES)
	go build -o $@

clean:
	rm -f rebed

test:
	go test ./...
