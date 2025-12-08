package metacomponent

import (
	"strings"

	"charm.land/lipgloss/v2"
)

var ()

func (m Model) View() string {
	if m.CurrentComponent == "" {
		return lipgloss.NewStyle().
			Width(m.width).
			Height(m.height).
			Align(lipgloss.Center, lipgloss.Center).
			Render("Select a component to view metadata")
	}

	var s strings.Builder

	s.WriteString(titleStyle.Render("Metadata: " + m.CurrentComponent))
	if m.Tick {
		s.WriteString(" " + m.Spinner.View())
	}
	s.WriteString("\n")

	s.WriteString(headerStyle.Render("Package:"))
	s.WriteString("\n")
	s.WriteString(itemStyle.Render(m.Metadata.PackageName))
	s.WriteString("\n")

	s.WriteString(headerStyle.Render("Children:"))
	s.WriteString("\n")
	if len(m.Metadata.Children) > 0 {
		for _, child := range m.Metadata.Children {
			s.WriteString(itemStyle.Render("- " + child))
			s.WriteString("\n")
		}
	} else {
		s.WriteString(itemStyle.Render("None"))
		s.WriteString("\n")
	}

	s.WriteString(headerStyle.Render("Usages:"))
	s.WriteString("\n")
	if len(m.Metadata.Usages) > 0 {
		for _, usage := range m.Metadata.Usages {
			s.WriteString(itemStyle.Render("- " + usage))
			s.WriteString("\n")
		}
	} else {
		s.WriteString(itemStyle.Render("None"))
		s.WriteString("\n")
	}

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Padding(1).
		Render(s.String())
}
