package componentlist

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	return lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(m.List.View())
}
