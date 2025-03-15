package models

import (
	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/tui/views"
)

type RestoreModel struct {
	Done     bool
	Quitting bool
}

func NewRestoreModel() RestoreModel {
	return RestoreModel{}
}

func (m RestoreModel) Init() tea.Cmd {
	return nil
}

func (m RestoreModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m RestoreModel) View() string {
	if m.Quitting {
		return "\n  Exiting installation...\n\n"
	}
	return views.RestoreView(m.Done)
}
