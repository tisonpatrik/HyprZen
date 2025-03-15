package components

import (
	"fmt"

	"hyprzen/internal/tui/styles"
)

func Checkbox(label string, checked bool) string {
	if checked {
		return styles.CheckboxStyle.Render("[x] " + label)
	}
	return fmt.Sprintf("[ ] %s", label)
}
