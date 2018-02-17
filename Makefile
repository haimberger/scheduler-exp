.PHONY: $(shell ls -d *)

default:
	@echo "Usage: make [command]"

test: clock.t

%.t:
	@go test -v -race --cover ./$* && echo ""

test-coverage: clock.cv

%.cv:
	@go test -coverprofile=$*/coverage.out ./$* && go tool cover -html=$*/coverage.out
