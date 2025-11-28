package componentlist

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/goferwplynie/bubbleWaffle/internal/creator"
	"github.com/goferwplynie/bubbleWaffle/internal/models"
)

type Model struct {
	List         list.Model
	Keys         KeyMap
	Width        int
	Height       int
	LastSelected string
}

func New() Model {
	components, err := creator.GetComponents(".")
	if err != nil {
		components = []string{}
	}
	var items []list.Item
	for _, v := range components {
		items = append(items, models.Component{
			Name: v,
		})
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false

	itemList := list.New(items, delegate, 0, 0)
	itemList.Title = "Components"
	itemList.SetShowHelp(false) // We will handle help globally

	return Model{
		List: itemList,
		Keys: DefaultKeyMap(),
	}
}
