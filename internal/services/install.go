package services

import (
	"github.com/charmbracelet/bubbletea"

	"hyprzen/internal/core"
)

type InstallDoneMsg struct{}

func InstallCmd() tea.Cmd {
	return func() tea.Msg {
		core.PreInstall()
		return InstallDoneMsg{}
	}
}
