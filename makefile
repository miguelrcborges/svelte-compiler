.PHONY: build install uninstall

SRC = $(wildcard *.go **/*.go)
BIN = svelte-compiler

build: $(BIN)

$(BIN): $(SRC)
	go build -ldflags='-s -w' .

install: build
	cp $(BIN) /usr/bin/$(BIN)

uninstall:
	rm /usr/bin/$(BIN)
