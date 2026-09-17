# Simple Command Runner (Go CLI)

A minimalist CLI utility written in Go designed to record, inspect, and execute shell commands sequentially. Created as a practical hands-on project to learn Go fundamentals, file I/O, argument parsing, and process execution on macOS.

---

## Features

- **register**: Append shell commands to a persistent text-based storage (`commands.txt`).
- **check**: List all recorded commands with 1-based index line numbers.
- **run +<line_number>**: Execute a specific recorded command directly via macOS shell (`/bin/zsh`).

---

## Project Structure

```text
.
├── main.go          # Core application logic
├── commands.txt     # Auto-generated storage for registered commands
└── README.md        # Documentation
