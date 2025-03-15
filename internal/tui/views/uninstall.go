package views

import (
	"fmt"

	"hyprzen/internal/tui/styles"
)

func UninstallView(done bool) string {
	label := "Uninstalling HyprZen..."
	if done {
		label = "Unnstallation complete! Press 'q' to exit."
	}
	return fmt.Sprintf(
		"%s\n\n%s",
		styles.KeywordStyle.Render(label),
		styles.KeywordStyle.Render("Preparing system..."),
	)
}
