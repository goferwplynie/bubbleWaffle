package componentlist

import (
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	var loading string
	if m.Loading {
		loading = "Loading components " + m.spinner.View() + "\n"
	}
	view := tea.NewView("")

	view.SetContent(lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height).
		Padding(1).
		Render(loading + m.List.View()))
	return view
}
