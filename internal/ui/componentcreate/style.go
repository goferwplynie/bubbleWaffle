package componentcreate

import "github.com/charmbracelet/lipgloss"

var (
	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF75B7")).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#FF75B7")).
			Padding(0, 1).
			Width(30)

	checkboxStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("212")).
			PaddingLeft(1)

	activeCheckboxStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("205")).
				PaddingLeft(1).
				Bold(true)

	buttonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFF7DB")).
			Background(lipgloss.Color("#FF75B7")).
			Padding(0, 3).
			MarginTop(1)

	activeButtonStyle = buttonStyle.
				Foreground(lipgloss.Color("#FFF7DB")).
				Background(lipgloss.Color("#C22A71")).
				Bold(true)

	popupStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("#874BFD")).
			Padding(1, 2).
			Width(50)
)
