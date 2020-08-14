.PHONY: all rollerderby install uninstall clean test test-all test-e2e
DEST = $(shell pwd)/build

# test runner (can be overridden by CI)
GOTEST ?= go test

all: rollerderby

rollerderby:
	@go build -o "$(DEST)/rollerderby" ./cmd/rollerderby
	@echo "$(DEST)/rollerderby"

install: all
	@cp -f $(DEST)/* $$GOPATH/bin/

uninstall:
	@rm -f $GOPATH/bin/rollerderby

clean:
	@rm -rf build/

test: test-all

test-all:
	@$(GOTEST) -v -count 1 `go list ./... | grep -v test/e2e`

test-e2e:
	@$(GOTEST) -v -count 1 `go list ./test/e2e` $(FLAGS)
