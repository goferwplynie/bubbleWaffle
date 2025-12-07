package metacomponent

import (
	"charm.land/bubbles/v2/spinner"
	"github.com/goferwplynie/bubbleWaffle/internal/models"
)

type Model struct {
	width, height    int
	CurrentComponent string
	Metadata         models.Metadata
	Spinner          spinner.Model
	Tick             bool
	CurrentPath      string
}

func New() Model {
	return Model{
		Spinner:     spinner.New(spinner.WithSpinner(spinner.Dot)),
		Tick:        false,
		CurrentPath: ".",
	}
}
