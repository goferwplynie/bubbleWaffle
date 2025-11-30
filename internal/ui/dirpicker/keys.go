package dirpicker

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	PickDir  key.Binding
	EnterDir key.Binding
	GoBack   key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		PickDir: key.NewBinding(
			key.WithKeys(" "),
			key.WithHelp("space", "pick directory"),
		),
		EnterDir: key.NewBinding(
			key.WithKeys("enter", "l", "right"),
			key.WithHelp("enter/l/→", "enter directory"),
		),
		GoBack: key.NewBinding(
			key.WithKeys("-", "h", "left"),
			key.WithHelp("-/h/←", "go back"),
		),
	}
}
