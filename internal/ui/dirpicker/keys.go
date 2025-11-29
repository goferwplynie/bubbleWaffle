package dirpicker

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	PickDir key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		PickDir: key.NewBinding(
			key.WithKeys("space"),
			key.WithHelp("space", "pick directory"),
		),
	}
}

