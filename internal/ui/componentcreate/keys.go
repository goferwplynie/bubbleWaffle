package componentcreate

import "charm.land/bubbles/v2/key"

type KeyMap struct {
	Up     key.Binding
	Down   key.Binding
	Enter  key.Binding
	Submit key.Binding
	Cancel key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "shift+tab"),
			key.WithHelp("↑/shift+tab", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "tab"),
			key.WithHelp("↓/tab", "down"),
		),
		Enter: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "select"),
		),
		Submit: key.NewBinding(
			key.WithKeys("shift+enter"),
			key.WithHelp("shift+enter", "submit"),
		),
		Cancel: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "cancel"),
		),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Up, k.Down, k.Enter, k.Submit, k.Cancel}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter, k.Submit, k.Cancel},
	}
}
