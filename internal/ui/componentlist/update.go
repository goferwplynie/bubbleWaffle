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
		m.Width, m.Height = msg.Width, msg.Height
		m.List.SetSize(msg.Width-6, msg.Height-4) // Reserve space for help
	}
	m.List, cmd = m.List.Update(msg)
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
