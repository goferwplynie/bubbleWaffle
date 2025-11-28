package componentlist

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/goferwplynie/bubbleWaffle/internal/creator"
	"github.com/goferwplynie/bubbleWaffle/internal/models"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width/4, msg.Height-3
		m.List.SetSize(msg.Width/2, msg.Height/2)
	}
	m.List, cmd = m.List.Update(msg)

	if m.List.SelectedItem() != nil {
		currentSelected := m.List.SelectedItem().FilterValue()
		if currentSelected != m.LastSelected {
			m.LastSelected = currentSelected
			return m, tea.Batch(cmd, func() tea.Msg {
				return models.ItemChangedMsg{Name: currentSelected}
			})
		}
	}

	return m, cmd
}

func (m *Model) RefreshList() {
	components, _ := creator.GetComponents(".")
	var items []list.Item
	for _, v := range components {
		items = append(items, models.Component{Name: v})
	}
	m.List.SetItems(items)
}
