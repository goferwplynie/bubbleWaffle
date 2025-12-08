package componentlist

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/goferwplynie/bubbleWaffle/internal/analyzer"
	"github.com/goferwplynie/bubbleWaffle/internal/models"
)

func (m Model) Init() tea.Cmd {
	return tea.Batch(LoadList, m.spinner.Tick)
}

func LoadList() tea.Msg {
	components, _ := analyzer.LoadComponents(".")
	var items []list.Item
	for _, v := range components {
		items = append(items, models.Component{
			Name: v.Name,
		})
	}
	return UpdateList{
		Items: items,
	}
}

type UpdateList struct {
	Items []list.Item
}
