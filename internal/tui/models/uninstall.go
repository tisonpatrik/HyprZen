package models

import (
	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/tui/views"
)

type UninstallModel struct {
	Done     bool
	Quitting bool
}

func NewUninstallModel() UninstallModel {
	return UninstallModel{}
}

func (m UninstallModel) Init() tea.Cmd {
	return nil
}

func (m UninstallModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m UninstallModel) View() string {
	if m.Quitting {
		return "\n  Exiting installation...\n\n"
	}
	return views.UninstallView(m.Done)
}
