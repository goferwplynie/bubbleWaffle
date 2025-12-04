package componentlist

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/goferwplynie/bubbleWaffle/internal/creator"
	"github.com/goferwplynie/bubbleWaffle/internal/models"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width/4, msg.Height-3
		m.List.SetSize(msg.Width/2, msg.Height/2)
	case updateList:
		cmd = m.List.SetItems(msg.Items)
		cmds = append(cmds, cmd)
		m.Loading = false
	case spinner.TickMsg:
		if m.Loading {
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}
	case ComponentCreatedMsg:
		m.Loading = true
		cmds = append(cmds, LoadList, m.spinner.Tick)

	}
	m.List, cmd = m.List.Update(msg)
	cmds = append(cmds, cmd)

	if m.List.SelectedItem() != nil {
		currentSelected := m.List.SelectedItem().FilterValue()
		if currentSelected != m.LastSelected {
			m.LastSelected = currentSelected
			cmd = func() tea.Msg {
				return models.ItemChangedMsg{Name: currentSelected}
			}
			cmds = append(cmds, cmd)
			return m, tea.Batch(cmds...)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) RefreshList(path string) {
	components, _ := creator.GetComponents(path)
	var items []list.Item
	for _, v := range components {
		items = append(items, models.Component{Name: v})
	}
	m.List.SetItems(items)
}

type ComponentCreatedMsg struct{}
