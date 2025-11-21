package compositor

import "github.com/charmbracelet/lipgloss"

func (m *Model) View() string {
	var layout string
	listView := listStyle.Render(m.list.View())
	layout += lipgloss.NewStyle().Align(lipgloss.Center, lipgloss.Center).Render(listView)
	return layout
}

