NAME := badrobot
GITHUB_ORG := controlplaneio
DOCKER_HUB_ORG := controlplane

### github.com/controlplaneio/ensure-content.git makefile-header START ###
ifeq ($(NAME),)
  $(error NAME required, please add "NAME := project-name" to top of Makefile)
else ifeq ($(GITHUB_ORG),)
    $(error GITHUB_ORG required, please add "GITHUB_ORG := controlplaneio" to top of Makefile)
else ifeq ($(DOCKER_HUB_ORG),)
    $(error DOCKER_HUB_ORG required, please add "DOCKER_HUB_ORG := controlplane" to top of Makefile)
endif

PKG := github.com/$(GITHUB_ORG)/$(NAME)
CONTAINER_REGISTRY_FQDN ?= docker.io
CONTAINER_REGISTRY_URL := $(CONTAINER_REGISTRY_FQDN)/$(DOCKER_HUB_ORG)/$(NAME)

SHELL := /bin/bash
BUILD_DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

GIT_MESSAGE := $(shell git -c log.showSignature=false \
	log --max-count=1 --pretty=format:"%H" 2>/dev/null)
GIT_SHA := $(shell git -c log.showSignature=false rev-parse HEAD 2>/dev/null)
GIT_TAG := $(shell bash -c 'TAG=$$(git -c log.showSignature=false \
	describe --tags --exact-match --abbrev=0 $(GIT_SHA) 2>/dev/null); echo "$${TAG:-dev}"')
GIT_UNTRACKED_CHANGES := $(shell git -c log.showSignature=false \
	status --porcelain 2>/dev/null)

ifneq ($(GIT_UNTRACKED_CHANGES),)
  GIT_COMMIT := $(GIT_COMMIT)-dirty
  ifneq ($(GIT_TAG),dev)
    GIT_TAG := $(GIT_TAG)-dirty
  endif
endif

CONTAINER_TAG ?= $(GIT_TAG)
CONTAINER_TAG_LATEST := $(CONTAINER_TAG)
CONTAINER_NAME := $(CONTAINER_REGISTRY_URL):$(CONTAINER_TAG)

# if no untracked changes and tag is not dev, release `latest` tag
ifeq ($(GIT_UNTRACKED_CHANGES),)
  ifneq ($(GIT_TAG),dev)
    CONTAINER_TAG_LATEST = latest
  endif
endif

CONTAINER_NAME_LATEST := $(CONTAINER_REGISTRY_URL):$(CONTAINER_TAG_LATEST)

# golang buildtime, more at https://github.com/jessfraz/pepper/blob/master/Makefile
CTIMEVAR=-X $(PKG)/version.GITCOMMIT=$(GITCOMMIT) -X $(PKG)/version.VERSION=$(VERSION)
GO_LDFLAGS=-ldflags "-w $(CTIMEVAR)"
GO_LDFLAGS_STATIC=-ldflags "-w $(CTIMEVAR) -extldflags -static"

export NAME CONTAINER_REGISTRY_URL BUILD_DATE GIT_MESSAGE GIT_SHA GIT_TAG \
  CONTAINER_TAG CONTAINER_NAME CONTAINER_NAME_LATEST CONTAINER_NAME_TESTING
### github.com/controlplaneio/ensure-content.git makefile-header END ###

LDFLAGS=-s -w \
        -X github.com/controlplaneio/badrobot/cmd.version=$(GIT_TAG)\
        -X github.com/controlplaneio/badrobot/cmd.commit=$(GIT_SHA)

PACKAGE = none
BATS_PARALLEL_JOBS := $(shell command -v parallel 2>&1 >/dev/null && echo '--jobs 20')

.PHONY: all
all: help

# ---

.PHONY: all
lint:
	@echo "+ $@"
	make lint-go-fmt

.PHONY: lint-go-fmt
lint-go-fmt: ## golang fmt check
	@echo "+ $@"
	gofmt -l -s ./pkg | grep ".*\.go"; if [ "$$?" = "0" ]; then exit 1; fi

# ---
.PHONY: test
test: ## unit and local acceptance tests
	@echo "+ $@"
	make test-unit build test-acceptance

.PHONY: check-and-reinit-submodules
check-and-reinit-submodules:
	@if git submodule status | grep "^[-+]" ; then \
		git submodule update --init; \
	fi

.PHONY: uninit-submodules
uninit-submodules:
	git submodule deinit -f .

.PHONY: test-acceptance
test-acceptance: check-and-reinit-submodules build ## acceptance tests
	@echo "+ $@"
	bash -xc 'cd test && ./bin/bats/bin/bats $(BATS_PARALLEL_JOBS) .'

.PHONY: test-unit
test-unit: ## golang unit tests
	@echo "+ $@"
	CGO_ENABLED=1 go test -race $$(go list ./... | grep -v '/vendor/') -run "$${RUN:-.*}"

.PHONY: test-unit-verbose
test-unit-verbose: ## golang unit tests (verbose)
	@echo "+ $@"
	CGO_ENABLED=1 go test -race -v $$(go list ./... | grep -v '/vendor/') -run "$${RUN:-.*}"

# ---

.PHONY: build
build: ## golang build
	@echo "+ $@"
	go build -ldflags "$(LDFLAGS)" -o ./dist/badrobot .

.PHONY: prune
prune: ## golang dependency prune
	@echo "+ $@"
	go mod tidy

# ---

.PHONY: help
help: ## parse jobs and descriptions from this Makefile
	@grep -E '^[ a-zA-Z0-9_-]+:([^=]|$$)' $(MAKEFILE_LIST) \
    | grep -Ev '^help\b[[:space:]]*:' \
    | sort \
    | awk 'BEGIN {FS = ":.*?##"}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

