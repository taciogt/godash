GOTEST ?= go test

.PHONY: setup
setup:
	go get -u github.com/rakyll/gotest
	go install github.com/rakyll/gotest

.PHONY: test
test:
	$(GOTEST) -v ./...