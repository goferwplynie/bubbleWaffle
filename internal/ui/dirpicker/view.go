package dirpicker

import "github.com/charmbracelet/lipgloss"

func (m Model) View() string {
	return lipgloss.NewStyle().
		Width(m.Width).
		Height(m.Height).
		Border(lipgloss.RoundedBorder()).
		Render("Enter into your ui folder:\n\n" + m.Fp.View())
}
