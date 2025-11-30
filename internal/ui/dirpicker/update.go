package dirpicker

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	m.Fp, cmd = m.Fp.Update(msg)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width/3, msg.Height/2
		m.Fp.SetHeight(msg.Height / 2)
	case tea.KeyMsg:
		if key.Matches(msg, m.Keys.PickDir) {
			return m, nil
		}
	}
	return m, cmd
}
