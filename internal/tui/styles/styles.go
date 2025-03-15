package styles

import (
	"github.com/charmbracelet/lipgloss"

	"hyprzen/internal/utils"
)

const (
	ProgressBarWidth  = 71
	ProgressFullChar  = "█"
	ProgressEmptyChar = "░"
	DotChar           = " • "
)

var (
	KeywordStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	SubtleStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	TicksStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("79"))
	CheckboxStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212"))
	ProgressEmpty = SubtleStyle.Render(ProgressEmptyChar)
	DotStyle      = lipgloss.NewStyle().
			Foreground(lipgloss.Color("236")).
			Render(DotChar)
	MainStyle = lipgloss.NewStyle().MarginLeft(2)
	Ramp      = utils.MakeRampStyles("#B14FFF", "#00FFA3", ProgressBarWidth)
)
