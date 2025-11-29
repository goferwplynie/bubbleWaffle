package dirpicker

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width/3, msg.Height/2
		m.Fp.SetHeight(m.Height)
	}
	var cmd tea.Cmd
	m.Fp, cmd = m.Fp.Update(msg)
	return m, cmd
}
