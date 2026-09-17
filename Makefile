.PHONY: install build clean

VERSION ?= v0.1.0
LDFLAGS = -s -w -X main.appVersion=$(VERSION)

# Install binary directly into $HOME/go/bin
install:
	go install -C core-operator -ldflags "$(LDFLAGS)" ./cmd/wallop

# Compile binary locally into ./bin/wallop
build:
	go build -C core-operator -ldflags "$(LDFLAGS)" -o ../bin/wallop ./cmd/wallop

# Remove local build artifacts
clean:
	rm -rf bin/ dist/