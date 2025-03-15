package views

import (
	"fmt"

	"hyprzen/internal/tui/styles"
)

func InstallView(done bool) string {
	label := "Installing HyprZen...\n\nPlease wait..."
	return fmt.Sprintf(
		"%s\n\n%s",
		styles.KeywordStyle.Render(label),
		styles.KeywordStyle.Render("System setup in progress..."),
	)
}
