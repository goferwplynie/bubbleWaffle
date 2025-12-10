package componentlist

import (
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/spinner"
)

type Model struct {
	List         list.Model
	Keys         KeyMap
	Width        int
	Height       int
	LastSelected string
	spinner      spinner.Model
	Loading      bool
	CurrentPath  string
}

func New() Model {
	var items []list.Item
	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = false

	itemList := list.New(items, delegate, 0, 0)
	itemList.Title = "Components"
	itemList.SetShowHelp(false)

	return Model{
		List:        itemList,
		Keys:        DefaultKeyMap(),
		spinner:     spinner.New(spinner.WithSpinner(spinner.Dot)),
		Loading:     true,
		CurrentPath: ".",
	}
}
