package compositor

import "github.com/charmbracelet/bubbles/list"

type Model struct {
	componentList *list.Model
}

func New() *Model {
	return &Model{}
}
