package utils

import (
	"os"
	"os/exec"
	"strings"
)

func RunCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(output)), err
}

func IsPackageInstalled(pkg string) bool {
	_, err := RunCommand("pacman", "-Q", pkg)
	return err == nil
}

func IsGrubDetected() bool {
	if !IsPackageInstalled("grub") {
		return false
	}

	if _, err := os.Stat("/boot/grub/grub.cfg"); os.IsNotExist(err) {
		return false
	}

	return true
}

func IsNvidiaDetected() bool {
	cmd := exec.Command("sh", "-c", `lspci -nn | grep -Ei "vga|3d|display"`)
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return strings.Contains(strings.ToLower(string(output)), "nvidia")
}
