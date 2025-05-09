package aur_helpers

import (
	"fmt"
	"hyprzen/internal/utils"
	"os"
)

const (
	yayRepo = "https://aur.archlinux.org/yay-bin.git"
	yayDir  = "/tmp/yay-build"
)

func InstallYay() {
	fmt.Println("[INFO] Installing yay AUR helper...")

	if yayExists() {
		fmt.Println("[INFO] yay already installed, skipping...")
		return
	}

	if !installDependencies() {
		return
	}

	if !cloneYayRepo() {
		return
	}

	if !buildAndInstallYay() {
		return
	}

	fmt.Println("[INFO] yay successfully installed.")
	cleanupYayBuild()

	initializeYay()
}

func yayExists() bool {
	_, err := utils.RunCommand("which", "yay")
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

func cloneYayRepo() bool {
	_ = os.RemoveAll(yayDir)

	_, err := utils.RunCommand("git", "clone", yayRepo, yayDir)
	if err != nil {
		fmt.Println("[ERROR] Failed to clone yay repository:", err)
		return false
	}
	return true
}

func buildAndInstallYay() bool {
	if err := os.Chdir(yayDir); err != nil {
		fmt.Println("[ERROR] Failed to enter yay directory:", err)
		return false
	}

	_, err := utils.RunCommand("makepkg", "-si", "--noconfirm")
	if err != nil {
		fmt.Println("[ERROR] Failed to build and install yay:", err)
		return false
	}
	return true
}

func cleanupYayBuild() {
	if err := os.RemoveAll(yayDir); err != nil {
		fmt.Println("[WARN] Could not clean up temporary directory:", err)
	}
}

func initializeYay() {
	fmt.Println("[INFO] Initializing yay development package tracking...")

	if _, err := utils.RunCommand("yay", "-Y", "--gendb"); err != nil {
		fmt.Println("[WARN] Failed to generate yay development DB:", err)
		return
	}

	if _, err := utils.RunCommand("yay", "-Y", "--devel", "--save"); err != nil {
		fmt.Println("[WARN] Failed to enable dev package updates:", err)
		return
	}

	fmt.Println("[INFO] yay development package tracking initialized.")
}
