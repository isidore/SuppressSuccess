package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]

	message := "Success"
	if len(args) >= 2 && args[0] == "--message" {
		message = args[1]
		args = args[2:]
	}

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "usage: suppress_success [--message <text>] <command> [args...]")
		os.Exit(1)
	}

	cmd := exec.Command(args[0], args[1:]...)
	var buf bytes.Buffer
	cmd.Stdout = &buf
	cmd.Stderr = &buf

	err := cmd.Run()

	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			fmt.Fprintf(os.Stderr, "error running command: %v\n", err)
			os.Exit(1)
		}
	}

	output, code := suppress(exitCode, message, buf.Bytes())
	os.Stdout.Write(output)
	os.Exit(code)
}

func suppress(exitCode int, message string, input []byte) ([]byte, int) {
	if exitCode != 0 {
		return input, exitCode
	}
	return []byte(message + "\n"), 0
}
