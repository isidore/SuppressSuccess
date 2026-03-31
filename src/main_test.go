package main

import (
	"testing"
)

func TestSuppress_success(t *testing.T) {
	output, code := suppress(0, "Success", []byte("some output\n"))
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	if string(output) != "Success\n" {
		t.Errorf("expected 'Success\\n', got %q", string(output))
	}
}

func TestSuppress_successCustomMessage(t *testing.T) {
	output, code := suppress(0, "Build complete", []byte("some output\n"))
	if code != 0 {
		t.Errorf("expected exit code 0, got %d", code)
	}
	if string(output) != "Build complete\n" {
		t.Errorf("expected 'Build complete\\n', got %q", string(output))
	}
}

func TestSuppress_failure(t *testing.T) {
	input := []byte("error: build failed\n")
	output, code := suppress(1, "Success", input)
	if code != 1 {
		t.Errorf("expected exit code 1, got %d", code)
	}
	if string(output) != string(input) {
		t.Errorf("expected original output to pass through, got %q", string(output))
	}
}

func TestSuppress_failurePreservesExitCode(t *testing.T) {
	_, code := suppress(42, "Success", []byte("output\n"))
	if code != 42 {
		t.Errorf("expected exit code 42, got %d", code)
	}
}
