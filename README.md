# Wallop

A lightweight command orchestrator and automation runner written in Go, designed to register, inspect, and execute multi-service shell tasks on macOS.

---

## Architecture Overview

```text
.
├── LICENSE
├── README.md
├── commands.txt                     # Persistent command registry
├── bin/
│   └── wallop                       # Compiled operator binary
├── core-operator/                   # Go operator module
│   ├── go.mod
│   ├── cmd/
│   │   └── wallop/
│   │       └── main.go              # CLI entry point
│   └── internal/
└── demo/                            # Integration demo gateways
    ├── calendar-gateway/
    │   ├── calendar.py
    │   └── sample_payload.json
    ├── openclaw-gateway/
    │   ├── openclaw_runner.py
    │   └── sample_payload.json
    └── system-sentinel/
        ├── sentinel.py
        └── sample_payload.json

```

---

## Prerequisites

* macOS with `zsh` default shell
* Go 1.18+
* Python 3.8+ (standard library only)

---

## Build & Installation

From the project root, compile the binary directly into the root `bin/` directory:

```bash
# Build binary using the core-operator module
go build -C core-operator -o ../bin/wallop ./cmd/wallop

# Grant execution permissions to demo scripts
chmod +x demo/calendar-gateway/calendar.py \
         demo/openclaw-gateway/openclaw_runner.py \
         demo/system-sentinel/sentinel.py

```

> **Note:** Alternatively, `cd core-operator && go build -o ../bin/wallop ./cmd/wallop` achieves the same result.

---

## CLI Commands

| Command | Usage | Description |
| --- | --- | --- |
| `register` | `./bin/wallop register <command>` | Appends a shell command string to `commands.txt`. |
| `check` | `./bin/wallop check` | Lists all registered commands with line numbers. |
| `run` | `./bin/wallop run +<index>` | Executes the command at `<index>` via `/bin/zsh -c`. |

---

## Registering Demo Payloads

Each gateway reads input from a payload file, JSON string, or standard input. Register them directly from the project root:

### 1. System Sentinel

Inspects system and hardware metrics.

```bash
./bin/wallop register python3 demo/system-sentinel/sentinel.py demo/system-sentinel/sample_payload.json

```

### 2. Calendar Gateway

Handles calendar event synchronization and creation.

```bash
./bin/wallop register python3 demo/calendar-gateway/calendar.py demo/calendar-gateway/sample_payload.json

```

### 3. OpenClaw Gateway

Performs bounded HTTP requests and parses web text.

```bash
./bin/wallop register python3 demo/openclaw-gateway/openclaw_runner.py demo/openclaw-gateway/sample_payload.json

```

---

## Quickstart Walkthrough

Once registered, inspect and trigger the pipelines from the project root:

```bash
# 1. View all registered pipelines
./bin/wallop check

# 2. Run Sentinel check
./bin/wallop run +1

# 3. Run Calendar integration
./bin/wallop run +2

# 4. Run OpenClaw scraper
./bin/wallop run +3

```

---

## Clean Up

To reset local state and build artifacts:

```bash
rm -f commands.txt bin/wallop

```
