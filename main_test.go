package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestGreetingOutput(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected string
	}{
		{"default", nil, "Hello, world, from fabrik-test-beta"},
		{"with argument", []string{"Alice"}, "Hello, Alice, from fabrik-test-beta"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdArgs := append([]string{"run", "."}, tt.args...)
			out, err := exec.Command("go", cmdArgs...).CombinedOutput()
			if err != nil {
				t.Fatalf("go run failed: %v\nOutput: %s", err, out)
			}
			got := strings.TrimSpace(string(out))
			if got != tt.expected {
				t.Errorf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
