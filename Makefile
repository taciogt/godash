GOTEST ?= go test
GOTOOL ?= go tool

default: help

# TODO: colorful help, but needs improvements on formatting
.PHONY: help
help: # Show help for each of the Makefile recipes.
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done

#help: ## Show this help.
#	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

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

doc: setup  # Generate the project docs and make them available at http://localhost:6060/pkg/github.com/taciogt/godash
	godoc -http=:6060