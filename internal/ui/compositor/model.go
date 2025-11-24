package compositor

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/goferwplynie/bubbleWaffle/internal/ui/componentcreate"
	"github.com/goferwplynie/bubbleWaffle/internal/ui/componentlist"
	zone "github.com/lrstanley/bubblezone"
)

const (
	ListView = iota
	CreateView
)

type Model struct {
	List   componentlist.Model
	Create componentcreate.Model
	Help   help.Model
	State  int
	Width  int
	Height int
}

func New() *Model {
	return &Model{
		List:   componentlist.New(),
		Create: componentcreate.New(),
		Help:   help.New(),
		State:  ListView,
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(m.List.Init(), m.Create.Init())
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width, m.Height = msg.Width, msg.Height
		m.Help.Width = msg.Width
		// Update children size
		m.List, cmd = m.List.Update(msg)
		cmds = append(cmds, cmd)
		m.Create, cmd = m.Create.Update(msg)
		cmds = append(cmds, cmd)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

		switch m.State {
		case ListView:
			if key.Matches(msg, m.List.Keys.NewComponent) {
				m.State = CreateView
				// Reset create model state
				m.Create = componentcreate.New()
				m.Create.Width, m.Create.Height = m.Width, m.Height
				return m, nil
			}
			m.List, cmd = m.List.Update(msg)
			cmds = append(cmds, cmd)

		case CreateView:
			if key.Matches(msg, m.Create.Keys.Cancel) {
				m.State = ListView
				return m, nil
			}
			m.Create, cmd = m.Create.Update(msg)
			cmds = append(cmds, cmd)
		}

	case componentcreate.ComponentCreatedMsg:
		m.List.RefreshList()
		m.State = ListView
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	var content string
	var helpView string

	switch m.State {
	case ListView:
		content = m.List.View()
		helpView = m.Help.View(m.List.Keys)
	case CreateView:
		content = m.Create.View()
		helpView = m.Help.View(m.Create.Keys)
	}

	return zone.Scan(content + "\n" + helpView)
}
