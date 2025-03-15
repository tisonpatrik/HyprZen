package models

import (
	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/services"
	"hyprzen/internal/tui/views"
)

type InstallModel struct {
	Done     bool
	Quitting bool
}

func NewInstallModel() InstallModel {
	return InstallModel{}
}

// Init spustí instalaci hned po startu modelu
func (m InstallModel) Init() tea.Cmd {
	return services.InstallCmd()
}

// Update zpracovává zprávy (dokončení instalace, ukončení aplikace)
func (m InstallModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case services.InstallDoneMsg:
		m.Done = true
		return m, tea.Quit // Automaticky ukončí aplikaci po instalaci

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		}
	}
	return m, nil
}

// View zobrazí stav instalace
func (m InstallModel) View() string {
	if m.Quitting {
		return "\n  Exiting installation...\n\n"
	}
	return views.InstallView(m.Done)
}
