package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"

	"hyprzen/internal/tui"
)

func main() {
	p := tea.NewProgram(tui.NewMainModel())
	if _, err := p.Run(); err != nil {
		fmt.Println("could not start program:", err)
	}
}
