package componentlist

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	return lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height).
		Padding(1).
		Render(m.List.View())
}
