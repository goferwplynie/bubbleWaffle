package componentcreate

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/goferwplynie/bubbleWaffle/internal/creator"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.Keys.Down):
			m.SelectedOption++
			if m.SelectedOption > 4 {
				m.SelectedOption = 0
			}
		case key.Matches(msg, m.Keys.Up):
			m.SelectedOption--
			if m.SelectedOption < 0 {
				m.SelectedOption = 4
			}
		case key.Matches(msg, m.Keys.Enter):
			switch m.SelectedOption {
			case 1:
				m.CreateOptions.StyleFile.checked = !m.CreateOptions.StyleFile.checked
			case 2:
				m.CreateOptions.KeybindsFile.checked = !m.CreateOptions.KeybindsFile.checked
			case 3:
				m.CreateOptions.BubbleZone.checked = !m.CreateOptions.BubbleZone.checked
			case 4:
				opts := &creator.ComponentOptions{
					StyleFile:    m.CreateOptions.StyleFile.checked,
					KeybindsFile: m.CreateOptions.KeybindsFile.checked,
					BubbleZone:   m.CreateOptions.BubbleZone.checked,
				}
				if err := creator.CreateComponent(".", m.TextInput.Value(), opts); err != nil {
					m.Err = err
					return m, nil
				}
				return m, func() tea.Msg { return ComponentCreatedMsg{} }
			}
		}
	}

	if m.SelectedOption == 0 {
		m.TextInput.Focus()
		m.TextInput, cmd = m.TextInput.Update(msg)
		cmds = append(cmds, cmd)
	} else {
		m.TextInput.Blur()
	}

	return m, tea.Batch(cmds...)
}

type ComponentCreatedMsg struct{}
