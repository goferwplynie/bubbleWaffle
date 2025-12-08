package componentlist

import (
	lipgloss "charm.land/lipgloss/v2"
)

func (m Model) View() string {
	var loading string
	if m.Loading {
		loading = "Loading components " + m.spinner.View() + "\n"
	}

	return lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height).
		Padding(1).
		Render(loading + m.List.View())
}
