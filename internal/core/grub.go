package core

import (
	"fmt"
	"regexp"
	"strings"

	"hyprzen/internal/utils"
)

func ConfigureGrub() {
	if !utils.IsGrubDetected() {
		fmt.Println("[INFO] GRUB není detekován, přeskočeno...")
		return
	}

	fmt.Println("[INFO] GRUB detekován, začíná konfigurace...")

	if _, err := utils.RunCommand("sudo", "cp", "/etc/default/grub", "/etc/default/grub.hyprzen.bkp"); err != nil {
		fmt.Println("[ERROR] Nelze zálohovat /etc/default/grub:", err)
		return
	}
	if _, err := utils.RunCommand("sudo", "cp", "/boot/grub/grub.cfg", "/boot/grub/grub.hyprzen.bkp"); err != nil {
		fmt.Println("[ERROR] Nelze zálohovat /boot/grub/grub.cfg:", err)
		return
	}

	if utils.IsNvidiaDetected() {
		setupNVIDIA()
	}

	fmt.Println("[INFO] Stahuji a instaluji Catppuccin GRUB téma...")
	_, _ = utils.RunCommand(
		"git", "clone", "https://github.com/catppuccin/grub.git",
	)

	if _, err := utils.RunCommand("sudo", "cp", "-r", "grub/src/catppuccin-mocha-grub-theme", "/boot/grub/themes"); err != nil {
		fmt.Println("[ERROR] Nelze zkopírovat Catppuccin GRUB téma:", err)
		return
	}

	grubThemePath := "/boot/grub/themes/catppuccin-mocha-grub-theme/theme.txt"

	if _, err := utils.RunCommand("sudo", "sed", "-i", fmt.Sprintf(`/^GRUB_THEME=/c\GRUB_THEME="%s"`, grubThemePath), "/etc/default/grub"); err != nil {
		fmt.Println("[ERROR] Nelze nastavit GRUB téma:", err)
		return
	}

	if _, err := utils.RunCommand("sudo", "grub-mkconfig", "-o", "/boot/grub/grub.cfg"); err != nil {
		fmt.Println("[ERROR] Nelze vygenerovat nový GRUB config:", err)
		return
	}
	if _, err := utils.RunCommand("rm", "-rf", "grub"); err != nil {
		fmt.Println("[ERROR] Nelze odstranit adresar grub", err)
	}

	fmt.Println("[INFO] GRUB konfigurace dokončena.")
}

func setupNVIDIA() {
	fmt.Println(
		"[INFO] Nvidia GPU detekována, přidávám `nvidia_drm.modeset=1`...",
	)
	grubConfig, _ := utils.RunCommand(
		"grep",
		"^GRUB_CMDLINE_LINUX_DEFAULT=",
		"/etc/default/grub",
	)

	if grubConfig == "" {
		fmt.Println(
			"[INFO] GRUB_CMDLINE_LINUX_DEFAULT nenalezen, vytvářím nový...",
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
	if _, err := utils.RunCommand("sudo", "sed", "-i", sedCmd, "/etc/default/grub"); err != nil {
		fmt.Println("[ERROR] Nelze upravit /etc/default/grub:", err)
		return
	}
}
