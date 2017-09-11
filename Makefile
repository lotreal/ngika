all: build

.PHONY: build
build:
	go build -v -ldflags=-s -o build/ngika
