package componentlist

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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
	return updateList{
		Items: items,
	}
}

type updateList struct {
	Items []list.Item
}
