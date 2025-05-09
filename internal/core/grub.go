package core

import (
	"fmt"
	"regexp"
	"strings"

	"hyprzen/internal/utils"
)

const (
	// GRUB configuration paths
	grubDefaultSrc = "/etc/default/grub"
	grubDefaultBkp = "/etc/default/grub.hyprzen.bkp"
	grubCfgSrc     = "/boot/grub/grub.cfg"
	grubCfgBkp     = "/boot/grub/grub.hyprzen.bkp"

	// Theme paths
	catppuccinThemePath = "/boot/grub/themes/catppuccin-mocha-grub-theme"
	catppuccinThemeBkp  = "/boot/grub/themes/catppuccin-mocha-grub-theme.hyprzen.bkp"
	grubThemePath       = "/boot/grub/themes/catppuccin-mocha-grub-theme/theme.txt"
	clonedThemePath     = "grub/src/catppuccin-mocha-grub-theme"

	// Repository
	gitRepo = "https://github.com/catppuccin/grub.git"
)

func ConfigureGrub(nvidiaIsDetected bool) {
	if !utils.IsGrubDetected() {
		return
	}
	makeGrubBackup()
	if nvidiaIsDetected {
		setupNVIDIA()
	}
	setupTheme()
	cleanup()
}

func setupTheme() {
	_, _ = utils.RunCommand(
		"git", "clone", gitRepo,
	)

	if _, err := utils.RunCommand("sudo", "cp", "-r", clonedThemePath, catppuccinThemePath); err != nil {
		fmt.Println("[ERROR] Cannot copy Catppuccin GRUB theme:", err)
		return
	}

	if _, err := utils.RunCommand("sudo", "sed", "-i", fmt.Sprintf(`/^GRUB_THEME=/c\GRUB_THEME="%s"`, grubThemePath), grubDefaultSrc); err != nil {
		fmt.Println("[ERROR] Cannot set GRUB theme:", err)
		return
	}

	if _, err := utils.RunCommand("sudo", "grub-mkconfig", "-o", grubCfgSrc); err != nil {
		fmt.Println("[ERROR] Cannot generate new GRUB config:", err)
		return
	}
}

func makeGrubBackup() {
	if _, err := utils.RunCommand("sudo", "cp", grubDefaultSrc, grubDefaultBkp); err != nil {
		fmt.Println("[ERROR] Cannot backup", grubDefaultSrc+":", err)
		return
	}
	if _, err := utils.RunCommand("sudo", "cp", grubCfgSrc, grubCfgBkp); err != nil {
		fmt.Println("[ERROR] Cannot backup", grubCfgSrc+":", err)
		return
	}
}

func setupNVIDIA() {
	grubConfig, _ := utils.RunCommand(
		"grep",
		"^GRUB_CMDLINE_LINUX_DEFAULT=",
		grubDefaultSrc,
	)
	if grubConfig == "" {
		fmt.Println(
			"[INFO] GRUB_CMDLINE_LINUX_DEFAULT not found, creating new...",
		)
		grubConfig = `GRUB_CMDLINE_LINUX_DEFAULT="quiet splash"`
	}

	re := regexp.MustCompile(`\bnvidia_drm\.modeset=1\b`)
	grubConfig = re.ReplaceAllString(grubConfig, "")

	reValue := regexp.MustCompile(`GRUB_CMDLINE_LINUX_DEFAULT="(.*?)"`)
	matches := reValue.FindStringSubmatch(grubConfig)

	var newValue string
	if len(matches) > 1 {
		existingParams := matches[1]
		newValue = fmt.Sprintf(
			`GRUB_CMDLINE_LINUX_DEFAULT="%s nvidia_drm.modeset=1"`,
			strings.TrimSpace(existingParams),
		)
	} else {
		newValue = `GRUB_CMDLINE_LINUX_DEFAULT="quiet splash nvidia_drm.modeset=1"`
	}
	sedCmd := fmt.Sprintf(`/^GRUB_CMDLINE_LINUX_DEFAULT=/c\%s`, newValue)
	if _, err := utils.RunCommand("sudo", "sed", "-i", sedCmd, grubDefaultSrc); err != nil {
		fmt.Println("[ERROR] Cannot edit", grubDefaultSrc+":", err)
		return
	}
}

func cleanup() {
	if _, err := utils.RunCommand("rm", "-rf", "grub"); err != nil {
		fmt.Println("[ERROR] Cannot delete grub directory:", err)
	}
}
