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
		expected []string
	}{
		{"default", nil, []string{"Hello, world, from fabrik-test-beta", "e2e-cross-repo-spawn"}},
		{"with argument", []string{"Alice"}, []string{"Hello, Alice, from fabrik-test-beta", "e2e-cross-repo-spawn"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmdArgs := append([]string{"run", "."}, tt.args...)
			out, err := exec.Command("go", cmdArgs...).CombinedOutput()
			if err != nil {
				t.Fatalf("go run failed: %v\nOutput: %s", err, out)
			}
			got := string(out)
			for _, s := range tt.expected {
				if !strings.Contains(got, s+"\n") {
					t.Errorf("expected output to contain line %q, got %q", s, got)
				}
			}
		})
	}
}
