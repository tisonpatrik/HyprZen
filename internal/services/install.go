package services

import (
	"github.com/charmbracelet/bubbletea"

	"hyprzen/internal/core"
)

type InstallDoneMsg struct{}

// InstallCmd spustí instalaci a čeká na dokončení
func InstallCmd() tea.Cmd {
	return func() tea.Msg {
		core.Install()          // Spustíme instalaci a počkáme na její dokončení
		return InstallDoneMsg{} // Vrátíme zprávu až po dokončení instalace
	}
}
