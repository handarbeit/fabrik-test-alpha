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
		{
			"default",
			nil,
			[]string{
				"Hello, world, from fabrik-test-beta",
				"e2e-cross-repo-spawn",
				"e2e-cross-repo-spawn-31",
				"e2e-cross-repo-spawn-52",
				"e2e-cross-repo-spawn-3472",
				"e2e-cross-repo-spawn-3496",
				"e2e-cross-repo-spawn-3620",
				"e2e-cross-repo-spawn-3737",
				"e2e-cross-repo-spawn-3834",
				"e2e-cross-repo-spawn-3855",
				"e2e-cross-repo-spawn-3922",
				"e2e-cross-repo-spawn-20260801-125714-6717",
				"e2e-cross-repo-spawn-20260801-144217-6757",
				"e2e-cross-repo-spawn-20260803-043620-7283",
				"e2e-cross-repo-spawn-20260804-114237-6848",
				"e2e-cross-repo-spawn-20260810-015521-8654",
				"e2e-cross-repo-spawn-20260810-042952-3676",
				"e2e-cross-repo-spawn-20260813-045054-8185",
				"e2e-cross-repo-spawn-20260813-132237-0922",
				"e2e-cross-repo-spawn-20260813-200250-3163",
				"e2e-cross-repo-spawn-20260814-033648-2864",
				"e2e-cross-repo-spawn-20260814-052928-1443",
			},
		},
		{
			"with argument",
			[]string{"Alice"},
			[]string{
				"Hello, Alice, from fabrik-test-beta",
				"e2e-cross-repo-spawn",
				"e2e-cross-repo-spawn-31",
				"e2e-cross-repo-spawn-52",
				"e2e-cross-repo-spawn-3472",
				"e2e-cross-repo-spawn-3496",
				"e2e-cross-repo-spawn-3620",
				"e2e-cross-repo-spawn-3737",
				"e2e-cross-repo-spawn-3834",
				"e2e-cross-repo-spawn-3855",
				"e2e-cross-repo-spawn-3922",
				"e2e-cross-repo-spawn-20260801-125714-6717",
				"e2e-cross-repo-spawn-20260801-144217-6757",
				"e2e-cross-repo-spawn-20260803-043620-7283",
				"e2e-cross-repo-spawn-20260804-114237-6848",
				"e2e-cross-repo-spawn-20260810-015521-8654",
				"e2e-cross-repo-spawn-20260810-042952-3676",
				"e2e-cross-repo-spawn-20260813-045054-8185",
				"e2e-cross-repo-spawn-20260813-132237-0922",
				"e2e-cross-repo-spawn-20260813-200250-3163",
				"e2e-cross-repo-spawn-20260814-033648-2864",
				"e2e-cross-repo-spawn-20260814-052928-1443",
			},
		},
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
