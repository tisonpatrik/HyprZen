package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/tui/components"
	"hyprzen/internal/tui/styles"
)

type RestoreView struct {
	Progress float64
	Done     bool
}

func NewRestoreView() RestoreView {
	return RestoreView{}
}

func (m RestoreView) Init() tea.Cmd {
	return nil
}

func (m RestoreView) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case tea.KeyMsg:
		return NewInstallView(), nil
	}
	return m, nil
}

func (m RestoreView) View() string {
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
