package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/tui/views"
)

type MainModel struct {
	Choice   int
	Quitting bool
}

func NewMainModel() MainModel {
	return MainModel{}
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice = (m.Choice + 1) % 4
		case "k", "up":
			m.Choice = (m.Choice - 1 + 4) % 4
		case "enter":
			switch m.Choice {
			case 0:
				return views.NewInstallView(), nil
			case 1:
				return views.NewRestoreView(), nil
			case 2:
				return views.NewUninstallView(), nil
			case 3:
				m.Quitting = true
				return m, tea.Quit
			}
		}
	}

	return m, nil
}

func (m MainModel) View() string {
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	return views.MainView(m.Choice)
}
