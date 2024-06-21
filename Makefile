GOTEST ?= go test  ## Optional custom go test tool
GOTOOL ?= go tool

default: help

.PHONY: help
help:  ## Show this help
	./help.sh "$(MAKEFILE_LIST)"

.PHONY: setup
setup:
	go get -u github.com/rakyll/gotest
	go install github.com/rakyll/gotest
	go get golang.org/x/tools/cmd/godoc
	go install golang.org/x/tools/cmd/godoc

.PHONY: test
test:
	$(GOTEST) -v ./...

COVERAGE_OUT = coverage.out
$(COVERAGE_OUT): *.go
	$(GOTEST) -cover -coverprofile=coverage.out -v ./...

coverage-report: $(COVERAGE_OUT)
	$(GOTOOL) cover -html=coverage.out

doc: setup  ## Generate the project docs and make them available at http://localhost:6060/pkg/github.com/taciogt/godash
	godoc -http=:6060