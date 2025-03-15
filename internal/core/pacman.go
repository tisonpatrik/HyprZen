package core

import (
	"fmt"
	"os"

	"hyprzen/internal/utils"
)

// Konfiguruje pacman.conf
func ConfigurePacman() {
	if _, err := os.Stat("/etc/pacman.conf.hyde.bkp"); err == nil {
		fmt.Println("[INFO] Pacman již nakonfigurován, přeskočeno...")
		return
	}

	fmt.Println("[INFO] Konfigurace pacman.conf...")

	_, _ = utils.RunCommand(
		"sudo", "cp", "/etc/pacman.conf", "/etc/pacman.conf.hyde.bkp",
	)

	// Úprava pacman.conf
	_, _ = utils.RunCommand(
		"sudo",
		"sed",
		"-i",
		"/^#Color/c\\Color\nILoveCandy",
		"/etc/pacman.conf",
	)
	_, _ = utils.RunCommand(
		"sudo",
		"sed",
		"-i",
		"/^#VerbosePkgLists/c\\VerbosePkgLists",
		"/etc/pacman.conf",
	)
	_, _ = utils.RunCommand(
		"sudo",
		"sed",
		"-i",
		"/^#ParallelDownloads/c\\ParallelDownloads = 5",
		"/etc/pacman.conf",
	)
	_, _ = utils.RunCommand("sudo", "pacman", "-Syyu")

	fmt.Println("[INFO] Pacman konfigurace dokončena.")
}
