package internal

import (
	"fmt"
	"strings"
)

// RunPreInstall sets up system configs before installing Hyprland.
func RunPreInstall() {
	fmt.Println("🛠 Running pre-install configuration...")

	configureBootloader()
	configurePacman()

	fmt.Println("✅ Pre-install setup complete.")
}

// configureBootloader sets up GRUB or systemd-boot
func configureBootloader() {
	if IsPackageInstalled("grub") && FileExists("/boot/grub/grub.cfg") {
		fmt.Println("🔹 Configuring GRUB...")

		if !FileExists("/etc/default/grub.hyzen.bkp") {
			RunCommand("sudo cp /etc/default/grub /etc/default/grub.hyzen.bkp")
			RunCommand("sudo cp /boot/grub/grub.cfg /boot/grub/grub.hyzen.bkp")

			if DetectNvidia() {
				fmt.Println(
					"🎮 Nvidia detected, adding `nvidia_drm.modeset=1` to GRUB.",
				)
				RunCommand(
					`sudo sed -i '/^GRUB_CMDLINE_LINUX_DEFAULT=/c\GRUB_CMDLINE_LINUX_DEFAULT="quiet splash nvidia_drm.modeset=1"' /etc/default/grub`,
				)
			}

			RunCommand("sudo grub-mkconfig -o /boot/grub/grub.cfg")
		} else {
			fmt.Println("✔️ GRUB is already configured.")
		}
	}

	if IsPackageInstalled("systemd") && DetectNvidia() && IsSystemdBoot() {
		fmt.Println("🔹 Configuring systemd-boot...")

		entries, _ := RunCommand(
			"find /boot/loader/entries/ -type f -name '*.conf'",
		)
		for file := range strings.SplitSeq(entries, "\n") {
			if FileExists(file) && !FileExists(file+".hyzen.bkp") {
				RunCommand("sudo cp " + file + " " + file + ".hyzen.bkp")
				RunCommand(
					`sudo sed -i '/^options/c\options quiet splash nvidia_drm.modeset=1' ` + file,
				)
			}
		}
	}
}

// configurePacman modifies pacman settings
func configurePacman() {
	if FileExists("/etc/pacman.conf") &&
		!FileExists("/etc/pacman.conf.hyzen.bkp") {
		fmt.Println("📦 Configuring pacman...")

		RunCommand("sudo cp /etc/pacman.conf /etc/pacman.conf.hyzen.bkp")
		RunCommand(
			`sudo sed -i '/^#Color/c\Color\nILoveCandy' /etc/pacman.conf`,
		)
		RunCommand(
			`sudo sed -i '/^#VerbosePkgLists/c\VerbosePkgLists' /etc/pacman.conf`,
		)
		RunCommand(
			`sudo sed -i '/^#ParallelDownloads/c\ParallelDownloads = 5' /etc/pacman.conf`,
		)
		RunCommand(`sudo sed -i '/^#\[multilib\]/,+1 s/^#//' /etc/pacman.conf`)

		RunCommand("sudo pacman -Syyu --noconfirm")
		RunCommand("sudo pacman -Fy")
	} else {
		fmt.Println("✔️ Pacman is already configured.")
	}
}
