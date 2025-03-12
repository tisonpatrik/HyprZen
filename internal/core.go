package internal

import (
	"fmt"
	"os"
	"os/exec"
)

// Check if running as root
func CheckRoot() {
	if os.Getuid() == 0 {
		fmt.Println("❌ HyZen should not be run as root.")
		os.Exit(1)
	}
}

// Run a system command and return output
func RunCommand(cmd string) string {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		return fmt.Sprintf("❌ Error: %s\n%s", err, out)
	}
	return string(out)
}
