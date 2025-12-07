package utils

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func ViewToString(view tea.View) string {
	return lipgloss.NewCanvas(lipgloss.NewLayer(view.Content)).Render()
}
