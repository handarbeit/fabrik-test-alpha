package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestGreetingOutput(t *testing.T) {
	out, err := exec.Command("go", "run", ".").Output()
	if err != nil {
		t.Fatalf("go run . failed: %v", err)
	}
	got := strings.TrimSpace(string(out))
	if !strings.Contains(got, "from fabrik-test-beta") {
		t.Errorf("expected output to contain %q, got %q", "from fabrik-test-beta", got)
	}
	if got != "Hello, world, from fabrik-test-beta" {
		t.Errorf("expected %q, got %q", "Hello, world, from fabrik-test-beta", got)
	}
}
