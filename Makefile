# vi: ft=make

GOPATH := $(shell go env GOPATH)
BREW_PREFIX := $(shell brew --prefix)

# Get names of the tools from tools.go
TOOLS = $(shell cat tools.go | grep _ | awk -F'"' '{print $$2}' | sed -E 's~@(latest|v[0-9.]*)~~g' | awk -F'/' '{print $$NF}')

.PHONY: help
help:  ## print this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort -d | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# install binaries defined in tools.go
$(addprefix $(GOPATH)/bin/,$(TOOLS)):
	@echo Installing $(notdir $@) from tools.go
	@cat tools.go | grep _ | grep "$(notdir $@)" | awk -F'"' '{print $$2}' | xargs -tI % go install %

.PHONY: setup-tools
setup-tools: $(addprefix $(GOPATH)/bin/,$(TOOLS)) ## installs all tools defined in tools.go

$(BREW_PREFIX)/bin/golangci-lint:
	brew install golangci-lint

$(BREW_PREFIX)/bin/pre-commit:
	brew install pre-commit

.git/hooks/pre-commit:
	pre-commit install --install-hooks

.PHONY: setup-git-hooks
setup-git-hooks: $(BREW_PREFIX)/bin/pre-commit .git/hooks/pre-commit

.PHONY: test
test: $(GOPATH)/bin/gotest ## run the tests
	gotest -p 1 -v -count=1 ./...

.PHONY: fmt
fmt: $(GOPATH)/bin/goimports ## format go code
	goimports -w .

.PHONY: lint
lint: $(BREW_PREFIX)/bin/golangci-lint ## lint go code
	golangci-lint run --fix