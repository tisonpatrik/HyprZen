package views

import (
	"hyprzen/internal/tui/styles"
)

func InstallView(done bool) string {
	if done {
		return styles.KeywordStyle.Render(
			"Installation complete!\n\nPress 'q' to exit.",
		)
	}

	return styles.KeywordStyle.Render("Installing HyprZen...\n\nPlease wait...")
}
