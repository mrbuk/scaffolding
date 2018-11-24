all: test lint 
.PHONY : all

test: 
	go test ./...

lint:
	golint .
