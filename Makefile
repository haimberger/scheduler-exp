.PHONY: $(shell ls -d *)

default:
	@echo "Usage: make [command]"

lint: clock.lint task.lint

%.lint:
	@command -v gometalinter || (go get -u github.com/alecthomas/gometalinter && gometalinter --install)
	@gometalinter ./$*

test: clock.test task.test

%.test:
	@go test -v -race --cover ./$* && echo ""

test-coverage: clock.cov task.cov

%.cov:
	@go test -coverprofile=$*/coverage.out ./$* && go tool cover -html=$*/coverage.out

test-update: task.update

%.update:
	@go test -v --cover ./$* -update && echo ""
