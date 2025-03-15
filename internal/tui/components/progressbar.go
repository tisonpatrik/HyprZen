package components

import (
	"fmt"
	"math"
	"strings"

	"hyprzen/internal/tui/styles"
)

func Progressbar(percent float64) string {
	w := float64(styles.ProgressBarWidth)

	fullSize := int(math.Round(w * percent))
	fullCells := strings.Builder{}

	for _, style := range styles.Ramp[:fullSize] {
		fullCells.WriteString(style.Render(styles.ProgressFullChar))
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(styles.ProgressEmpty, emptySize)

	return fmt.Sprintf(
		"%s%s %3.0f%%",
		fullCells.String(),
		emptyCells,
		math.Round(percent*100),
	)
}
