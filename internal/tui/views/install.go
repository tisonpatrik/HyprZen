package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/tui/components"
	"hyprzen/internal/tui/styles"
)

type InstallView struct {
	Progress float64
	Done     bool
}

func NewInstallView() InstallView {
	return InstallView{}
}

func (m InstallView) Init() tea.Cmd {
	return nil
}

func (m InstallView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return NewInstallView(), nil
	}
	return m, nil
}

func (m InstallView) View() string {
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
