package tui

import (
	"fmt"

	"hyprzen/internal/tui/styles"
	"hyprzen/internal/tui/views"
)

func (m Model) View() string {
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	if !m.Chosen {
		return m.choicesView()
	}
	return m.chosenView()
}

func (m Model) choicesView() string {
	c := m.Choice

	tpl := "Welcome to HyprZen\n\n"
	tpl += "%s\n\n"
	tpl += styles.SubtleStyle.Render("j/k, up/down: select") + styles.DotStyle +
		styles.SubtleStyle.Render("enter: choose") + styles.DotStyle +
		styles.SubtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		views.Checkbox("install", c == 0),
		views.Checkbox("restore", c == 1),
		views.Checkbox("uninstall", c == 2),
		views.Checkbox("quit", c == 3),
	)

	return fmt.Sprintf(tpl, choices)
}

func (m Model) chosenView() string {
	var msg string

	switch m.Choice {
	case 0:
		msg = fmt.Sprintf("Installing...\n\nPreparing %s and %s...",
			styles.KeywordStyle.Render("hyprland"),
			styles.KeywordStyle.Render("wayland-utils"),
		)
	case 1:
		msg = fmt.Sprintf(
			"Restoring previous configuration...\n\nFetching %s...",
			styles.KeywordStyle.Render("backup"),
		)
	case 2:
		msg = fmt.Sprintf("Uninstalling...\n\nRemoving %s and %s...",
			styles.KeywordStyle.Render("hyprland"),
			styles.KeywordStyle.Render("configs"),
		)
	default:
		msg = fmt.Sprintf("Simply lovely....\n\nMaybe something %s can %s...",
			styles.KeywordStyle.Render("nice"),
			styles.KeywordStyle.Render("happen"),
		)
	}

	label := "Processing..."
	if m.Loaded {
		label = "Done! Press 'q' to exit."
	}

	return msg + "\n\n" + label + "\n" + views.Progressbar(m.Progress) + "%"
}
