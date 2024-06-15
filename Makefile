GOTEST ?= go test
GOTOOL ?= go tool

.PHONY: setup
setup:
	go get -u github.com/rakyll/gotest
	go install github.com/rakyll/gotest

.PHONY: test
test:
	$(GOTEST) -v ./...

COVERAGE_OUT = coverage.out
$(COVERAGE_OUT): *.go
	$(GOTEST) -cover -coverprofile=coverage.out -v ./...

coverage-report: $(COVERAGE_OUT)
	$(GOTOOL) cover -html=coverage.out