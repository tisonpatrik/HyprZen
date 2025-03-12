package tests

import (
	"os"
	"os/exec"
	"testing"
)

// Test that HyZen should NOT be run as root
func TestCheckRoot(t *testing.T) {
	if os.Getuid() == 0 {
		t.Fatal(
			"❌ HyZen should not run as root, but test was executed as root!",
		)
	}
}

// Test if the application runs without crashing
func TestAppRuns(t *testing.T) {
	cmd := exec.Command("go", "run", "../cmd/hyzen/main.go") // Adjust path
	err := cmd.Run()
	if err != nil {
		t.Fatalf("❌ Failed to run main.go: %v", err)
	}
}
