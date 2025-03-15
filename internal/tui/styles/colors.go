package styles

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
)

func makeRampStyles(colorA, colorB string, steps float64) (s []lipgloss.Style) {
	cA, _ := colorful.Hex(colorA)
	cB, _ := colorful.Hex(colorB)

	for i := 0.0; i < steps; i++ {
		c := cA.BlendLuv(cB, i/steps)
		s = append(
			s,
			lipgloss.NewStyle().Foreground(lipgloss.Color(colorToHex(c))),
		)
	}
	return
}

func colorToHex(c colorful.Color) string {
	return fmt.Sprintf(
		"#%s%s%s",
		floatToHex(c.R),
		floatToHex(c.G),
		floatToHex(c.B),
	)
}

func floatToHex(f float64) string {
	s := strconv.FormatInt(int64(f*255), 16)
	if len(s) == 1 {
		s = "0" + s
	}
	return s
}
