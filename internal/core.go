package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Check if running as root
func CheckRoot() {
	if os.Getuid() == 0 {
		fmt.Println("❌ HyZen should not be run as root.")
		os.Exit(1)
	}
}

// RunCommand executes a shell command and returns output
func RunCommand(cmd string) (string, error) {
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	return string(out), err
}

// fileExists checks if a file exists
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// isPackageInstalled checks if a package is installed
func IsPackageInstalled(pkg string) bool {
	_, err := RunCommand("pacman -Q " + pkg)
	return err == nil
}

// detectNvidia checks if an Nvidia GPU is present
func DetectNvidia() bool {
	output, _ := RunCommand("lspci | grep -i nvidia")
	return strings.Contains(output, "NVIDIA")
}

// isSystemdBoot checks if systemd-boot is used
func IsSystemdBoot() bool {
	output, _ := RunCommand(
		"bootctl status 2>/dev/null | awk '{if ($1 == \"Product:\") print $2}'",
	)
	return strings.Contains(output, "systemd-boot")
}
