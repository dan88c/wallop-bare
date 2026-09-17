package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Injected at build time via -ldflags; defaults to "v0.1.0" for local dev
var appVersion = "v0.1.0"

const storeFile = "commands.txt"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:\n  register <command>\n  check\n  run +<line_number>\n  version")
		return
	}

	action := os.Args[1]

	switch action {
	case "version", "-v", "--version":
		fmt.Printf("core-operator version %s\n", appVersion)

	case "register":
		if len(os.Args) < 3 {
			fmt.Println("Usage: register <command text>")
			return
		}
		cmdText := strings.Join(os.Args[2:], " ")
		register(cmdText)

	case "check":
		check()

	case "run":
		if len(os.Args) < 3 || !strings.HasPrefix(os.Args[2], "+") {
			fmt.Println("Usage: run +<line_number>")
			return
		}
		lineStr := strings.TrimPrefix(os.Args[2], "+")
		lineNum, err := strconv.Atoi(lineStr)
		if err != nil || lineNum < 1 {
			fmt.Println("Error: invalid line number")
			return
		}
		runLine(lineNum)

	default:
		fmt.Println("Unknown command:", action)
	}
}

func register(entry string) {
	f, err := os.OpenFile(storeFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer f.Close()

	if _, err := f.WriteString(entry + "\n"); err != nil {
		fmt.Printf("Error writing entry: %v\n", err)
		return
	}
	fmt.Println("Registered successfully.")
}

func check() {
	lines, err := readLines()
	if err != nil {
		fmt.Printf("Error reading records: %v\n", err)
		return
	}

	if len(lines) == 0 {
		fmt.Println("No records found.")
		return
	}

	for idx, line := range lines {
		fmt.Printf("%d: %s\n", idx+1, line)
	}
}

func runLine(lineNum int) {
	lines, err := readLines()
	if err != nil {
		fmt.Printf("Error reading records: %v\n", err)
		return
	}

	if lineNum > len(lines) {
		fmt.Printf("Error: line %d does not exist (total lines: %d)\n", lineNum, len(lines))
		return
	}

	targetCmd := lines[lineNum-1]
	fmt.Printf("Executing [line %d]: %s\n", lineNum, targetCmd)

	cmd := exec.Command("/bin/zsh", "-c", targetCmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Run(); err != nil {
		fmt.Printf("Command exited with error: %v\n", err)
	}
}

func readLines() ([]string, error) {
	if _, err := os.Stat(storeFile); os.IsNotExist(err) {
		return []string{}, nil
	}

	file, err := os.Open(storeFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}