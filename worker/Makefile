APP = worker
DATE = $(shell date +%FT%T%Z)
BUILD_DIR = build/bin
GIT_VER=$(shell git rev-parse HEAD)
GO = go

LDFLAGS=-ldflags "-s -w \
									-X github.com/frzifus/load-operator/worker/pkg/version.hash=${GIT_VER} \
									-X github.com/frzifus/load-operator/worker/pkg/version.buildtimestamp=${DATE}"

define go-build
	GOOS=$(1) GOARCH=$(2) $(GO) build -trimpath ${LDFLAGS} -o ${BUILD_DIR}/${APP}-$(1)-$(2) -v cmd/${APP}/*.go
endef

all: amd64 arm

.PHONY: build_deps
build_deps:
	go install golang.org/x/lint/golint@latest

.PHONY: amd64
amd64:
	$(call go-build,linux,amd64)

.PHONY: arm
arm:
	$(call go-build,linux,arm)

.PHONY: lint
lint:
	golint -set_exit_status ./pkg/... ./cmd/...

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	-rm -f ${BUILD_DIR}/${BINARY}-*

.PHONY: distclean
distclean:
	rm -rf ./build

.PHONY: mrproper
mrproper: distclean
	git ls-files --others | xargs rm -rf
