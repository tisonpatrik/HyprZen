package models

import (
	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/tui/views"
)

type InstallModel struct {
	Done     bool
	Quitting bool
}

func NewInstallModel() InstallModel {
	return InstallModel{}
}

func (m InstallModel) Init() tea.Cmd {
	return nil
}

func (m InstallModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		default:
			m.Done = true
		}
	}

	return m, nil
}

func (m InstallModel) View() string {
	if m.Quitting {
		return "\n  Exiting installation...\n\n"
	}
	return views.InstallView(m.Done)
}
