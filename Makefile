GO ?= go
DETECTED_LIBNBD_VERSION = $(shell dpkg-query --showformat='$${Version}' -W libnbd-dev || echo "0.0.0-libnbd-not-found")

default: build

.PHONY: build
build: build-dependencies migration-manager
	$(GO) build -o ./bin/migration-managerd ./cmd/migration-managerd
	$(GO) build -o ./bin/migration-manager-worker ./cmd/migration-manager-worker

.PHONY: build-dependencies
build-dependencies:
	@if ! dpkg --compare-versions 1.20 "<=" ${DETECTED_LIBNBD_VERSION}; then \
		echo "Please install libnbd-dev with version >= 1.20"; \
		exit 1; \
	fi

.PHONY: migration-manager
migration-manager:
	mkdir -p ./bin/
	CGO_ENABLED=0 GOARCH=amd64 $(GO) build -o ./bin/migration-manager.linux.amd64 ./cmd/migration-manager
	CGO_ENABLED=0 GOARCH=arm64 $(GO) build -o ./bin/migration-manager.linux.arm64 ./cmd/migration-manager
	GOOS=darwin GOARCH=amd64 $(GO) build -o ./bin/migration-manager.macos.amd64 ./cmd/migration-manager
	GOOS=darwin GOARCH=arm64 $(GO) build -o ./bin/migration-manager.macos.arm64 ./cmd/migration-manager
	GOOS=windows GOARCH=amd64 $(GO) build -o ./bin/migration-manager.windows.amd64.exe ./cmd/migration-manager
	GOOS=windows GOARCH=arm64 $(GO) build -o ./bin/migration-manager.windows.arm64.exe ./cmd/migration-manager

.PHONY: build-all-packages
build-all-packages:
	$(GO) mod tidy
	$(GO) build ./...
	$(GO) test -c -o /dev/null ./...

.PHONY: test
test: build-dependencies
	$(GO) test ./... -v -cover

.PHONY: static-analysis
static-analysis: build-dependencies
ifeq ($(shell command -v go-licenses),)
	(cd / ; $(GO) install -v -x github.com/google/go-licenses@latest)
endif
ifeq ($(shell command -v golangci-lint),)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$($(GO) env GOPATH)/bin
endif
ifeq ($(shell command -v shellcheck),)
	echo "Please install shellcheck"
	exit 1
endif
	go-licenses check --disallowed_types=forbidden,unknown,restricted --ignore libguestfs.org/libnbd ./...
	shellcheck --shell sh internal/worker/scripts/*.sh
	golangci-lint run ./...
	run-parts $(shell run-parts -V >/dev/null 2>&1 && echo -n "--verbose --exit-on-error --regex '\.sh\$'") scripts/lint

.PHONY: clean
clean:
	rm -rf dist/ bin/

.PHONY: release-snapshot
release-snapshot:
ifeq ($(shell command -v goreleaser),)
	echo "Please install goreleaser"
	exit 1
endif
	goreleaser release --snapshot --clean

.PHONY: build-dev-container
build-dev-container:
	docker build -t migration-manager-dev ./.devcontainer/

DOCKER_RUN := docker run -i -v .:/home/vscode/src --mount source=migration_manager_devcontainer_goroot,target=/go,type=volume --mount source=migration_manager_devcontainer_cache,target=/home/vscode/.cache,type=volume -w /home/vscode/src -u 1000:1000 migration-manager-dev

.PHONY: docker-build
docker-build: build-dev-container
	${DOCKER_RUN} make build

.PHONY: docker-build-all-packages
docker-build-all-packages: build-dev-container
	${DOCKER_RUN} make build-all-packages

.PHONY: docker-test
docker-test: build-dev-container
	${DOCKER_RUN} make test

.PHONY: docker-static-analysis
docker-static-analysis: build-dev-container
	${DOCKER_RUN} make static-analysis

.PHONY: docker-release-snapshot
docker-release-snapshot: build-dev-container
	${DOCKER_RUN} make release-snapshot

.PHONY: enter-dev-container
enter-dev-container:
	@docker exec -it -w /workspaces/migration-manager ${USER}_migration_manager_devcontainer /bin/bash
