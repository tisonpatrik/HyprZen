package aur_helpers

import (
	"fmt"
	"hyprzen/internal/utils"
	"os"
)

const (
	paruRepo = "https://aur.archlinux.org/paru.git"
	tmpDir   = "/tmp/paru-build"
)

func InstallParu() {
	fmt.Println("[INFO] Installing paru AUR helper...")

	if paruExists() {
		fmt.Println("[INFO] paru already installed, skipping...")
		return
	}

	if !installDependencies() {
		return
	}

	if !cloneParuRepo() {
		return
	}

	if !buildAndInstallParu() {
		return
	}

	fmt.Println("[INFO] paru successfully installed.")
	cleanupParuBuild()
}

// paruExists checks if paru binary is already available.
func paruExists() bool {
	_, err := utils.RunCommand("which", "paru")
	return err == nil
}

func installDependencies() bool {
	_, err := utils.RunCommand("sudo", "pacman", "-S", "--needed", "base-devel", "git")
	if err != nil {
		fmt.Println("[ERROR] Failed to install base-devel and git:", err)
		return false
	}
	return true
}

func cloneParuRepo() bool {
	_ = os.RemoveAll(tmpDir)

	_, err := utils.RunCommand("git", "clone", paruRepo, tmpDir)
	if err != nil {
		fmt.Println("[ERROR] Failed to clone paru repository:", err)
		return false
	}
	return true
}

func buildAndInstallParu() bool {
	if err := os.Chdir(tmpDir); err != nil {
		fmt.Println("[ERROR] Failed to enter paru directory:", err)
		return false
	}

	_, err := utils.RunCommand("makepkg", "-si", "--noconfirm")
	if err != nil {
		fmt.Println("[ERROR] Failed to build and install paru:", err)
		return false
	}
	return true
}

func cleanupParuBuild() {
	if err := os.RemoveAll(tmpDir); err != nil {
		fmt.Println("[WARN] Could not clean up temporary directory:", err)
	}
}
