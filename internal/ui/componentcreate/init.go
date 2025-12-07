package componentcreate

import (
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}
