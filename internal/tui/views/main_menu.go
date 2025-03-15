package views

import (
	"fmt"

	"hyprzen/internal/tui/components"
	"hyprzen/internal/tui/styles"
)

func MainView(choice int) string {
	tpl := "Welcome to HyprZen\n\n"
	tpl += "%s\n\n"
	tpl += styles.SubtleStyle.Render("j/k, up/down: select") + styles.DotStyle +
		styles.SubtleStyle.Render("enter: choose") + styles.DotStyle +
		styles.SubtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		components.Checkbox("Install", choice == 0),
		components.Checkbox("Restore", choice == 1),
		components.Checkbox("Uninstall", choice == 2),
		components.Checkbox("Quit", choice == 3),
	)

	return fmt.Sprintf(tpl, choices)
}
