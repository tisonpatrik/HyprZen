package core

import (
	"fmt"
	"log"

	"hyprzen/internal/utils"
)

const (
	backupPath = "/etc/pacman.conf.hyprzen.bkp"
	src        = "/etc/pacman.conf"
	dest       = "/etc/pacman.conf.hyprzen.bkp"
)

func ConfigurePacman() {

	fmt.Println("[INFO] Configuring pacman.conf...")
	backupPacmanConfig()
	applyPacmanConfig()
	verifyPacmanConfig()

	fmt.Println("[INFO] Pacman configuration completed.")
}

func applyPacmanConfig() {
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
}

func backupPacmanConfig() {

	_, err := utils.RunCommand("sudo", "cp", src, dest)
	if err != nil {
		fmt.Println("[ERROR] Failed to create pacman.conf backup:", err)
	} else {
		fmt.Println("[INFO] Backup of pacman.conf created at", dest)
	}
}

func verifyPacmanConfig() {
	fmt.Println("[INFO] Verifying pacman configuration...")
	_, err := utils.RunCommand("pacman", "-Qi", "pacman")
	if err == nil {
		fmt.Println("[INFO] Pacman is working correctly.")

		if _, removeErr := utils.RunCommand("sudo", "rm", "-f", backupPath); removeErr != nil {
			fmt.Printf("[WARN] Could not remove backup file: %v\n", removeErr)
		} else {
			fmt.Println("[INFO] Backup removed:", backupPath)
		}
		return
	}

	fmt.Println("[ERROR] Pacman configuration seems broken. Restoring backup...")

	_, restoreErr := utils.RunCommand("sudo", "cp", backupPath, "/etc/pacman.conf")
	if restoreErr != nil {
		log.Fatalf("[FATAL] Failed to restore backup: %v", restoreErr)
	}

	log.Fatal("[FATAL] Pacman configuration invalid. Backup restored. Exiting.")
}
