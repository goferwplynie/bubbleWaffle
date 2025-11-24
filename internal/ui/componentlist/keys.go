package componentlist

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	NewComponent key.Binding
	Quit         key.Binding
}

func DefaultKeyMap() KeyMap {
	return KeyMap{
		NewComponent: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "new component"),
		),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+c", "q"),
			key.WithHelp("q/ctrl+c", "quit"),
		),
	}
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.NewComponent, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.NewComponent, k.Quit},
	}
}
