package compositor

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"github.com/goferwplynie/bubbleWaffle/internal/models"
	"github.com/goferwplynie/bubbleWaffle/internal/ui/componentcreate"
	"github.com/goferwplynie/bubbleWaffle/internal/ui/componentlist"
	"github.com/goferwplynie/bubbleWaffle/internal/ui/dirpicker"
	"github.com/goferwplynie/bubbleWaffle/internal/ui/metacomponent"
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
		m.Help.SetWidth(msg.Width)
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
		m.List, cmd = m.List.Update(componentlist.ComponentCreatedMsg{})
		cmds = append(cmds, cmd)
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
		cmd = func() tea.Msg {
			items := componentlist.RefreshList(msg.New)
			return componentlist.UpdateList{
				Items: items,
			}
		}
		m.Meta.SetCurrentPath(msg.New)
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

func (m *Model) View() tea.View {
	var content string
	var helpView string

	view := tea.NewView("")
	view.AltScreen = true

	switch m.CurrentView {
	case MainView:
		listView := m.List.View()
		helpView = m.Help.View(m.List.Keys)

		listView = listComponentStyle.Render(listView + "\n" + helpView)

		meta := metaComponentStyle.Render(m.Meta.View())
		fp := m.Fp.View()
		content = lipgloss.JoinHorizontal(lipgloss.Top, listView, meta)
		content = lipgloss.NewStyle().Width(m.Width).Height(m.Height).Render(content)
		if m.State == FilePicker {
			content = lipgloss.NewStyle().Faint(true).Render(content)
			fpLayer := lipgloss.NewLayer(fp)
			mainLayer := lipgloss.NewLayer(content)
			xCenter := (m.Width / 2) - (fpLayer.GetWidth() / 2)
			yCenter := (m.Height / 2) - (fpLayer.GetHeight() / 2)

			view.SetContent(lipgloss.NewCanvas(mainLayer, fpLayer.Z(1).X(xCenter).Y(yCenter)).Render())
		} else {
			view.SetContent(content)
		}
	case CreateView:
		content = m.Create.View()
		helpView = m.Help.View(m.Create.Keys)
		view.SetContent(content + "\n" + helpView)
	default:
		view.SetContent("this view is not handled ;c")

	}

	return view
}
