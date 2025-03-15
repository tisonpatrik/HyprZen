package utils

import (
	"bufio"
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
	cmd := exec.Command("lspci", "-k")
	output, err := cmd.StdoutPipe()
	if err != nil {
		return false
	}
	if err := cmd.Start(); err != nil {
		return false
	}

	scanner := bufio.NewScanner(output)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "VGA") || strings.Contains(line, "3D") {
			if strings.Contains(strings.ToLower(line), "nvidia") {
				return true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return false
	}

	_ = cmd.Wait()
	return false
}
