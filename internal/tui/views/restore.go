package views

import (
	"fmt"

	"hyprzen/internal/tui/styles"
)

func RestoreView(done bool) string {
	label := "Restoring HyprZen..."
	if done {
		label = "Restoration is complete! Press 'q' to exit."
	}
	return fmt.Sprintf(
		"%s\n\n%s",
		styles.KeywordStyle.Render(label),
		styles.KeywordStyle.Render("Preparing system..."),
	)
}
