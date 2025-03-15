package views

import (
	"fmt"

	"hyprzen/internal/tui/styles"
)

func InstallView(done bool) string {
	label := "Installing HyprZen..."
	if done {
		label = "Installation complete! Press 'q' to exit."
	}
	return fmt.Sprintf(
		"%s\n\n%s",
		styles.KeywordStyle.Render(label),
		styles.KeywordStyle.Render("Preparing system..."),
	)
}
