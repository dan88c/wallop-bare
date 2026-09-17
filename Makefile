GO_DIR   := core-operator
BIN_DIR  := bin
BIN_NAME := wallop
BIN      := $(BIN_DIR)/$(BIN_NAME)
PREFIX   ?= $(HOME)/.local

VERSION  ?= v0.1.0
LDFLAGS  := -s -w -X main.appVersion=$(VERSION)

.PHONY: all help build build-windows build-local download install test tidy clean

all: build

help:
	@echo "Wallop Targets:"
	@echo "  make build          Build native binary into $(BIN)"
	@echo "  make build-windows  Cross-compile Windows binary ($(BIN_DIR)/$(BIN_NAME).exe)"
	@echo "  make download       Fetch prebuilt binary into $(BIN_DIR)/ (no Go required)"
	@echo "  make install        Install native binary to $(PREFIX)/bin"
	@echo "  make test           Run all Go unit tests"
	@echo "  make tidy           Tidy Go modules"
	@echo "  make clean          Remove compiled binaries"

build: build-local

build-local:
	@mkdir -p $(BIN_DIR)
	cd $(GO_DIR) && \
	go build \
	  -ldflags="$(LDFLAGS)" \
	  -o ../$(BIN) ./cmd/wallop
	@echo "Built: $(BIN) ($(VERSION))"

build-windows:
	@mkdir -p $(BIN_DIR)
	cd $(GO_DIR) && \
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build \
	  -ldflags="$(LDFLAGS)" \
	  -o ../$(BIN_DIR)/$(BIN_NAME).exe ./cmd/wallop
	@echo "Built: $(BIN_DIR)/$(BIN_NAME).exe"

download:
	@sh scripts/download.sh

install: build-local
	@mkdir -p $(PREFIX)/bin
	cp $(BIN) $(PREFIX)/bin/$(BIN_NAME)
	@echo "Installed to $(PREFIX)/bin/$(BIN_NAME)"

test:
	cd $(GO_DIR) && go test -v ./...

tidy:
	cd $(GO_DIR) && go mod tidy

clean:
	rm -rf $(BIN_DIR)