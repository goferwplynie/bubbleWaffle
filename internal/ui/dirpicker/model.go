package dirpicker

import "github.com/charmbracelet/bubbles/filepicker"

type Model struct {
	Fp            filepicker.Model
	Width, Height int
	Keys          KeyMap
}

func New() Model {
	fp := filepicker.New()
	fp.ShowPermissions = false
	fp.ShowHidden = false
	fp.ShowSize = false
	fp.DirAllowed = true
	return Model{
		Fp:   fp,
		Keys: DefaultKeyMap(),
	}
}
