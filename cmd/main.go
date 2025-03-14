package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/tui"
)

func main() {
	initialModel := tui.NewModel()
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}
