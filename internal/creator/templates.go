package creator

var ModelTemplate = `package {{.Name}}

type Model struct {}

func New() *Model{
	return &Model{}
}
	`
var InitTemplate = `package {{.Name}}

import(
	tea "charm.land/bubbletea/v2"
)

func (m *Model) Init() tea.Cmd{
	return nil
}
	`
var UpdateTemplate = `package {{.Name}}

import(
	tea "charm.land/bubbletea/v2"
)

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	return m, nil
}
	`

var ViewTemplate = `package {{.Name}}

func (m *Model) View() string{
	var layout string
	return layout
}
	`

var BubbleZoneView = `package {{.Name}}

import(
	zone "github.com/lrstanley/bubblezone"
)

func (m *Model) View() string{
	var layout string
	return zone.Mark(layout)
}
	`

var StylesTemplate = `package {{.Name}}

import(
	"github.com/charmbracelet/lipgloss"
)

var(
	cuteStyle = lipgloss.NewStyle()
)
	`
var KeyBindsTemplate = `package {{.Name}}

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {}

func DefaultKeyMap() KeyMap {
	return KeyMap{}
}`
