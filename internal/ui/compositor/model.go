package compositor

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/goferwplynie/bubbleWaffle/internal/models"
	"github.com/goferwplynie/bubbleWaffle/internal/ui/componentcreate"
	"github.com/goferwplynie/bubbleWaffle/internal/ui/componentlist"
	"github.com/goferwplynie/bubbleWaffle/internal/ui/dirpicker"
	"github.com/goferwplynie/bubbleWaffle/internal/ui/metacomponent"
	overlay "github.com/rmhubbert/bubbletea-overlay"
)

type View = int
type State = int

const (
	MainView View = iota
	CreateView

	List State = iota
	FilePicker
)

type Model struct {
	List        componentlist.Model
	Create      componentcreate.Model
	Meta        metacomponent.Model
	Fp          dirpicker.Model
	Help        help.Model
	CurrentView View
	State       State
	Width       int
	Height      int
	CurrentPath string
}

func New() *Model {
	return &Model{
		List:        componentlist.New(),
		Create:      componentcreate.New(),
		Help:        help.New(),
		Fp:          dirpicker.New(),
		Meta:        metacomponent.New(),
		CurrentView: MainView,
		State:       List,
		CurrentPath: ".",
	}
}

func (m *Model) Init() tea.Cmd {
	return tea.Batch(m.List.Init(), m.Create.Init(), m.Meta.Init(), m.Fp.Init())
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
		m.Meta, cmd = m.Meta.Update(msg)
		cmds = append(cmds, cmd)
		m.Fp, cmd = m.Fp.Update(msg)
		cmds = append(cmds, cmd)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		}

		switch m.CurrentView {
		case MainView:
			if key.Matches(msg, key.NewBinding(key.WithKeys("d", "esc"))) {
				switch m.State {
				case List:
					m.State = FilePicker
				case FilePicker:
					m.State = List
				}
				return m, nil
			}
			switch m.State {
			case List:
				if key.Matches(msg, m.List.Keys.NewComponent) {
					m.CurrentView = CreateView
					m.Create = componentcreate.New()
					m.Create.Width, m.Create.Height = m.Width, m.Height
					return m, nil
				}
				m.List, cmd = m.List.Update(msg)
				cmds = append(cmds, cmd)
			case FilePicker:
				m.Fp, cmd = m.Fp.Update(msg)
				cmds = append(cmds, cmd)
			}

		case CreateView:
			if key.Matches(msg, m.Create.Keys.Cancel) {
				m.CurrentView = MainView
				return m, nil
			}
			m.Create, cmd = m.Create.Update(msg)
			cmds = append(cmds, cmd)
		}

	case componentcreate.ComponentCreatedMsg:
		m.List.RefreshList(m.CurrentPath)
		m.CurrentView = MainView

	case models.ItemChangedMsg, models.ComponentMetaMsg:
		m.Meta, cmd = m.Meta.Update(msg)
		cmds = append(cmds, cmd)
	case spinner.TickMsg:
		m.Meta, cmd = m.Meta.Update(msg)
		cmds = append(cmds, cmd)

		m.List, cmd = m.List.Update(msg)
		cmds = append(cmds, cmd)
	case dirpicker.DirChanged:
		m.CurrentPath = msg.New
		m.State = List
		//refresh list
		m.List.RefreshList(m.CurrentPath)
		//update meta
		m.Meta, cmd = m.Meta.Update(models.ItemChangedMsg{
			//get selected item name
			Name: m.List.List.SelectedItem().FilterValue(),
		})
		cmds = append(cmds, cmd)
	default:
		m.List, cmd = m.List.Update(msg)
		cmds = append(cmds, cmd)
		m.Meta, cmd = m.Meta.Update(msg)
		cmds = append(cmds, cmd)
		m.Fp, cmd = m.Fp.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) View() string {
	var content string
	var helpView string

	switch m.CurrentView {
	case MainView:
		content = m.List.View()
		helpView = m.Help.View(m.List.Keys)
		content = listComponentStyle.Render(content + "\n" + helpView)

		meta := metaComponentStyle.Render(m.Meta.View())
		fp := m.Fp.View()
		content = lipgloss.JoinHorizontal(lipgloss.Top, content, meta)
		content = lipgloss.NewStyle().Width(m.Width).Height(m.Height).Render(content)
		if m.State == FilePicker {
			content = lipgloss.NewStyle().Faint(true).Render(content)
			content = overlay.Composite(fp, content, overlay.Center, overlay.Center, 0, 0)
		}
		return content
	case CreateView:
		content = m.Create.View()
		helpView = m.Help.View(m.Create.Keys)
		return content + "\n" + helpView
	}

	return "this view is not handled ;c"
}
