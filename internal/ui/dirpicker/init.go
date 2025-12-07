package dirpicker

import (
	tea "charm.land/bubbletea/v2"
)

func (m Model) Init() tea.Cmd {
	return m.Fp.Init()
}
