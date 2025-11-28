package metacomponent

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/goferwplynie/bubbleWaffle/internal/models"
)

type Model struct {
	width, height    int
	CurrentComponent string
	Metadata         models.Metadata
	Spinner          spinner.Model
	Tick             bool
}

func New() Model {
	return Model{
		Spinner: spinner.New(spinner.WithSpinner(spinner.Dot)),
		Tick:    false,
	}
}
