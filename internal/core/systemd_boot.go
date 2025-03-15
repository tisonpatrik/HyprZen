package core

import (
	"bufio"
	"fmt"
	"os/exec"
	"strings"

	"hyprzen/internal/utils"
)

func ConfigureSystemdBoot() {
	if !utils.IsPackageInstalled("systemd") {
		fmt.Println("[INFO] systemd není nainstalován, přeskočeno...")
		return
	}

	status, _ := utils.RunCommand("bootctl", "status")
	if !strings.Contains(status, "systemd-boot") {
		fmt.Println("[INFO] systemd-boot není aktivní, přeskočeno...")
		return
	}

	fmt.Println("[INFO] systemd-boot detekován, začíná konfigurace...")

	if utils.IsNvidiaDetected() {
		fmt.Println(
			"[INFO] Nvidia GPU detekována, přidávám `nvidia_drm.modeset=1` do systemd-boot...",
		)

		cmd := exec.Command(
			"find",
			"/boot/loader/entries/",
			"-type",
			"f",
			"-name",
			"*.conf",
		)
		output, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("[ERROR] Nelze spustit `find`:", err)
			return
		}
		if err := cmd.Start(); err != nil {
			fmt.Println("[ERROR] Chyba při spuštění `find`:", err)
			return
		}

		scanner := bufio.NewScanner(output)
		for scanner.Scan() {
			file := scanner.Text()
			if file == "" {
				continue
			}
			if _, err := utils.RunCommand("sudo", "cp", file, file+".hyde.bkp"); err != nil {
				fmt.Println("[ERROR] Nelze zálohovat:", file)
				continue
			}

			if _, err := utils.RunCommand("sudo", "sed", "-i", "/^options/c\\options quiet splash nvidia_drm.modeset=1", file); err != nil {
				fmt.Println("[ERROR] Nelze upravit:", file)
			}
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("[ERROR] Chyba při čtení souborů `find`:", err)
		}

		_ = cmd.Wait()
	}

	fmt.Println("[INFO] systemd-boot konfigurace dokončena.")
}
