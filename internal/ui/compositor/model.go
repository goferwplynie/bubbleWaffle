package compositor

import (
	"github.com/charmbracelet/bubbles/list"
)

type Model struct {
	list          list.Model
	Width, Height int
}

func New() *Model {
	return &Model{
		list:   list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0),
		Width:  0,
		Height: 0,
	}
}
