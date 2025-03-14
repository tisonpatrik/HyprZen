package tui

import (
	"fmt"
	"math"
	"strings"
)

func (m Model) View() string {
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	if !m.Chosen {
		return choicesView(m)
	}
	return chosenView(m)
}

func choicesView(m Model) string {
	c := m.Choice

	tpl := "wellcome to HyprZen\n\n"
	tpl += "%s\n\n"
	tpl += subtleStyle.Render("j/k, up/down: select") + dotStyle +
		subtleStyle.Render("enter: choose") + dotStyle +
		subtleStyle.Render("q, esc: quit")

	choices := fmt.Sprintf(
		"%s\n%s\n%s\n%s",
		checkbox("install", c == 0),
		checkbox("restore", c == 1),
		checkbox("uninstall", c == 2),
		checkbox("quit", c == 3),
	)

	return fmt.Sprintf(tpl, choices)
}

func chosenView(m Model) string {
	var msg string

	switch m.Choice {
	case 0:
		msg = fmt.Sprintf(
			"Carrot planting?\n\nCool, we'll need %s and %s...",
			keywordStyle.Render("libgarden"),
			keywordStyle.Render("vegeutils"),
		)
	case 1:
		msg = fmt.Sprintf(
			"A trip to the market?\n\nOkay, then we should install %s and %s...",
			keywordStyle.Render("marketkit"),
			keywordStyle.Render("libshopping"),
		)
	case 2:
		msg = fmt.Sprintf(
			"Reading time?\n\nOkay, cool, then we’ll need a library. Yes, an %s.",
			keywordStyle.Render("actual library"),
		)
	default:
		msg = fmt.Sprintf(
			"Simply lovely....\n\nMaybe something %s can %s...",
			keywordStyle.Render("nice"),
			keywordStyle.Render("happen"),
		)
	}

	label := "Downloading..."
	if m.Loaded {
		label = "Downloaded. Press 'q' to exit."
	}

	return msg + "\n\n" + label + "\n" + progressbar(m.Progress) + "%"
}

func checkbox(label string, checked bool) string {
	if checked {
		return checkboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}

func progressbar(percent float64) string {
	w := float64(progressBarWidth)

	fullSize := int(math.Round(w * percent))
	fullCells := strings.Builder{}

	for _, style := range ramp[:fullSize] {
		fullCells.WriteString(style.Render(progressFullChar))
	}

	emptySize := int(w) - fullSize
	emptyCells := strings.Repeat(progressEmpty, emptySize)

	return fmt.Sprintf(
		"%s%s %3.0f",
		fullCells.String(),
		emptyCells,
		math.Round(percent*100),
	)
}
