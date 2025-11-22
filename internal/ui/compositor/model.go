package compositor

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/goferwplynie/bubbleWaffle/internal/creator"
	"github.com/goferwplynie/bubbleWaffle/internal/models"
)

type Model struct {
	list          list.Model
	Width, Height int
}

func New() *Model {
	components, err := creator.GetComponents(".")
	if err != nil {
		panic(err)
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
	itemList.Title = "your components cutie :3"

	return &Model{
		list:   itemList,
		Width:  0,
		Height: 0,
	}
}
