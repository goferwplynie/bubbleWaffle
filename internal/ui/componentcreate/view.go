package componentcreate

import (
	"fmt"

	lipgloss "charm.land/lipgloss/v2"
)

func (m Model) View() string {
	var content string

	// Title
	content += lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FAFAFA")).Render("Create New Component") + "\n\n"

	// Input
	content += "Name:\n"
	content += inputStyle.Render(m.TextInput.View()) + "\n\n"

	// Options
	options := []struct {
		label   string
		checked bool
	}{
		{"Style File", m.CreateOptions.StyleFile.checked},
		{"Keybinds File", m.CreateOptions.KeybindsFile.checked},
		{"BubbleZone", m.CreateOptions.BubbleZone.checked},
	}

	for i, opt := range options {
		cursor := " "
		style := checkboxStyle
		if m.SelectedOption == i+1 {
			cursor = ">"
			style = activeCheckboxStyle
		}

		checked := "[ ]"
		if opt.checked {
			checked = "[x]"
		}

		content += fmt.Sprintf("%s %s %s\n", cursor, style.Render(checked), opt.label)
	}

	// Button
	btnStyle := buttonStyle
	if m.SelectedOption == 4 {
		btnStyle = activeButtonStyle
	}
	content += "\n" + btnStyle.Render("Create")

	return lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(
			popupStyle.Render(content),
		)
}
