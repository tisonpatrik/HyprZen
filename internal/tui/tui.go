package tui

import (
	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/tui/models"
)

func NewTUI() tea.Model {
	return models.NewMainMenu()
}
