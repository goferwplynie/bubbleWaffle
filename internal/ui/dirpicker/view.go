package dirpicker

import (
	tea "charm.land/bubbletea/v2"
	lipgloss "charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	view := tea.NewView("")
	view.SetContent(lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height).
		Border(lipgloss.RoundedBorder()).
		Render("Enter into your ui folder:\n\n" + m.Fp.View()))
	return view
}
