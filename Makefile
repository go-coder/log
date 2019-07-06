all: build test

.PHONY: build
build:
	@scripts/go-build

.PHONY: test	
test:
	@echo " > testing file ..."

.PHONY: clean
clean:
	@echo " > cleaning file ..."
	rm -rf ./bin/