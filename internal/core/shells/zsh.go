package shells

import (
	"fmt"
	"hyprzen/internal/utils"
	"os"
	"strings"
)

func InstallZsh() {
	fmt.Println("[INFO] Setting zsh as default shell...")

	runInstallZsh()

	zshPath, err := utils.RunCommand("which", "zsh")
	if err != nil || zshPath == "" {
		fmt.Println("[ERROR] zsh binary not found after install.")
		return
	}

	addZshToShells(zshPath)
	setZshAsDefaultShell(zshPath)

	fmt.Printf("[INFO] Default shell set to zsh (%s)\n", zshPath)
}

func runInstallZsh() {
	if _, err := utils.RunCommand("yay", "-S", "--noconfirm", "--needed", "zsh"); err != nil {
		fmt.Println("[WARN] Failed to install zsh via yay, trying pacman fallback...")

		if _, err := utils.RunCommand("sudo", "pacman", "-S", "--noconfirm", "--needed", "zsh"); err != nil {
			fmt.Println("[ERROR] Failed to install zsh via yay and pacman:", err)
			return
		}
	}
}

func addZshToShells(zshPath string) {
	output, _ := utils.RunCommand("cat", "/etc/shells")
	if !strings.Contains(output, zshPath) {
		_, err := utils.RunCommand("sudo", "sh", "-c", fmt.Sprintf("echo %s >> /etc/shells", zshPath))
		if err != nil {
			fmt.Println("[ERROR] Could not add zsh to /etc/shells:", err)
			return
		}
		fmt.Println("[INFO] zsh added to /etc/shells")
	}
}

func setZshAsDefaultShell(zshPath string) {
	user := os.Getenv("USER")
	if user == "" {
		fmt.Println("[ERROR] USER environment variable not found")
		return
	}
	if _, err := utils.RunCommand("chsh", "-s", zshPath); err != nil {
		fmt.Println("[ERROR] Failed to change shell to zsh:", err)
		return
	}
}
