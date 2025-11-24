package componentcreate

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type Option struct {
	label   string
	checked bool
}

type CreateOptions struct {
	StyleFile    Option
	KeybindsFile Option
	BubbleZone   Option
}

type Model struct {
	TextInput      textinput.Model
	SelectedOption int
	CreateOptions  CreateOptions
	Keys           KeyMap
	Width          int
	Height         int
	Err            error
}

func New() Model {
	ti := textinput.New()
	ti.Placeholder = "Component Name"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return Model{
		TextInput:      ti,
		SelectedOption: 0,
		CreateOptions: CreateOptions{
			StyleFile: Option{
				label:   "style file",
				checked: false,
			},
			KeybindsFile: Option{
				label:   "keybinds file",
				checked: false,
			},
			BubbleZone: Option{
				label:   "BubbleZone",
				checked: false,
			},
		},
		Keys: DefaultKeyMap(),
	}
}
