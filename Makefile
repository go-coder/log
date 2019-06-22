
GOFILES=$(wildcard *.go)
BACKENDS=$(wildcard ./backend/*.go)
EXAMPLE=./example/main.go
GOBIN=./bin

all: build test

build:
	@echo " > formating file ..."
	go fmt $(GOFILES) 
	go fmt $(BACKENDS)
	golangci-lint run $(GOFILES) 
	golangci-lint run $(BACKENDS)
	@rm -rf $(GOBIN)
	@-mkdir $(GOBIN)
	go build -o $(patsubst %.go, $(GOBIN)/%, $(notdir $(EXAMPLE))) $(EXAMPLE)

test:
	@echo " > testing file ..."

.PHONY:clean
clean:
	@echo " > cleaning file ..."
	rm -rf $(GOBIN)