.PHONY: $(shell ls -d *)

default:
	@echo "Usage: make [command]"

test: clock.t task.t

%.t:
	@go test -v -race --cover ./$* && echo ""

test-coverage: clock.cv task.cv

%.cv:
	@go test -coverprofile=$*/coverage.out ./$* && go tool cover -html=$*/coverage.out

test-update: task.u

%.u:
	@go test -v --cover ./$* -update && echo ""
