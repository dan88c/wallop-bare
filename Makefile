.PHONY: install build test-demo clean

# Install binary to $HOME/go/bin for global access
install:
	go install -C core-operator ./cmd/wallop
	chmod +x demo/calendar-gateway/calendar.py demo/openclaw-gateway/openclaw_runner.py demo/system-sentinel/sentinel.py

# Build binary into ./bin/wallop
build:
	go build -C core-operator -o ../bin/wallop ./cmd/wallop
	chmod +x demo/calendar-gateway/calendar.py demo/openclaw-gateway/openclaw_runner.py demo/system-sentinel/sentinel.py

# Reset commands and build output
clean:
	rm -f commands.txt bin/wallop