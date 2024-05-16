TEMPLE_WATCH_DIR := ./internal/templates
GO_FILES := $(shell find . -name '*.go')

.PHONY: all clean build run templ-watch

all: build

clean:
	rm -rf ./bin

build:
	templ generate --dir $(TEMPLE_WATCH_DIR)
	go build -o ./target/bin/main main.go

dev:
	build
	air

templ-watch:
	templ generate --watch $(TEMPLE_WATCH_DIR) --proxy http://localhost:3000