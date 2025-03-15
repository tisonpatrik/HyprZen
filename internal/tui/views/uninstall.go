package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/tui/components"
	"hyprzen/internal/tui/styles"
)

type UninstallView struct {
	Progress float64
	Done     bool
}

func NewUninstallView() UninstallView {
	return UninstallView{}
}

func (m UninstallView) Init() tea.Cmd {
	return nil
}

func (m UninstallView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return NewUninstallView(), nil
	}
	return m, nil
}

func (m UninstallView) View() string {
	label := "Installing..."
	if m.Done {
		label = "Done! Press 'q' to exit."
	}

	return fmt.Sprintf(
		"Installing Hyprland...\n\n%s\n\n%s\n\n%s",
		styles.KeywordStyle.Render("Preparing system"),
		components.Progressbar(m.Progress),
		label,
	)
}
