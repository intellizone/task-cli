# Makefile for Task Tracker CLI

BINARY=task-cli
SRC=cmd/server/server.go

.PHONY: build run clean

build:
	go build -o $(BINARY) $(SRC)

run: build
	./$(BINARY) list

clean:
	rm -f $(BINARY)
