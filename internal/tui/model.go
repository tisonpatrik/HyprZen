package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fogleman/ease"
)

type Model struct {
	Choice   int
	Chosen   bool
	Frames   int
	Progress float64
	Loaded   bool
	Quitting bool
}

type (
	TickMsg  struct{}
	FrameMsg struct{}
)

func NewModel() Model {
	return Model{
		Choice:   0,
		Chosen:   false,
		Frames:   0,
		Progress: 0.0,
		Loaded:   false,
		Quitting: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

// Main update function.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if !m.Chosen {
		return updateChoices(msg, m)
	}
	return updateChosen(msg, m)
}

func updateChoices(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Choice++
			if m.Choice > 3 {
				m.Choice = 3
			}
		case "k", "up":
			m.Choice--
			if m.Choice < 0 {
				m.Choice = 0
			}
		case "enter":
			m.Chosen = true
			return m, frame()
		}
	}
	return m, nil
}

func updateChosen(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case frameMsg:
		if !m.Loaded {
			m.Frames++
			m.Progress = ease.OutBounce(float64(m.Frames) / float64(100))
			if m.Progress >= 1 {
				m.Progress = 1
				m.Loaded = true
			}
			return m, frame()
		}
	}

	return m, nil
}
