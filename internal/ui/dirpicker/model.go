package dirpicker

import (
	"path/filepath"

	"github.com/charmbracelet/bubbles/filepicker"
)

type Model struct {
	Fp            filepicker.Model
	Width, Height int
	Keys          KeyMap
}

func New() Model {
	fp := filepicker.New()
	fp.ShowPermissions = false
	fp.ShowHidden = true
	fp.ShowSize = false
	fp.DirAllowed = true
	fp.CurrentDirectory, _ = filepath.Abs("./")
	return Model{
		Fp:   fp,
		Keys: DefaultKeyMap(),
	}
}
