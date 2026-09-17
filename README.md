# Wallop

> [!NOTE]
> **Experimental Substrate for AI Agent Tool Execution**  
> This repository is a proof-of-concept (PoC) exploring a lightweight, native execution layer designed to interface with AI agents and orchestrate local shell/Python workflows.
> 
> * **Intent:** Testing execution patterns and tool-calling interfaces between AI reasoning loops and local system tasks.
> * **Status:** Sandbox prototype under active exploration; interfaces and CLI ergonomics are subject to rapid iteration.

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

## Installation

### Option 1: Pre-built Binaries (No Go Required)

Download the pre-compiled binary matching your architecture from [GitHub Releases](https://github.com/dan88c/wallop/releases) and install it into your system PATH:

#### For Apple Silicon Mac (M1/M2/M3/M4):
```bash
chmod +x wallop-darwin-arm64
sudo mv wallop-darwin-arm64 /usr/local/bin/wallop

```

#### For Intel Mac:

```bash
chmod +x wallop-darwin-amd64
sudo mv wallop-darwin-amd64 /usr/local/bin/wallop

```

#### For Linux (x86_64):

```bash
chmod +x wallop-linux-amd64
sudo mv wallop-linux-amd64 /usr/local/bin/wallop

```

Verify installation:

```bash
wallop version

```

---

### Option 2: Build from Source via Makefile (Requires Go)

If you clone the repository and prefer to compile locally with the same flags as the release build:

```bash
# Install binary globally to $HOME/go/bin
make install

# Or build locally to ./bin/wallop
make build

```

> **PATH Setup (One-time):** If using `make install`, ensure Go's bin directory is in your PATH:
> ```bash
> echo 'export PATH="$HOME/go/bin:$PATH"' >> ~/.zshrc && source ~/.zshrc
> 
> ```
> 
> 

To clean up build artifacts:

```bash
make clean
```

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

## Acknowledgements

Architected as an exploration into AI Agent tool execution layers. Implementation details, cross-compilation workflows, and Go scaffolding were developed using AI-assisted pair programming.