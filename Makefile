.PHONY: $(shell ls -d *)

default:
	@echo "Usage: make [command]"

lint: clock.lint

%.lint:
	@command -v gometalinter || (go get -u github.com/alecthomas/gometalinter && gometalinter --install)
	@gometalinter ./$*

test: clock.t

%.t:
	@go test -v -race --cover ./$* && echo ""

test-coverage: clock.cv

%.cv:
	@go test -coverprofile=$*/coverage.out ./$* && go tool cover -html=$*/coverage.out
