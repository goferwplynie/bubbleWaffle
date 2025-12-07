package metacomponent

import (
	"charm.land/bubbles/v2/spinner"
	tea "charm.land/bubbletea/v2"
	"github.com/goferwplynie/bubbleWaffle/internal/analyzer"
	"github.com/goferwplynie/bubbleWaffle/internal/models"
)

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width/4, msg.Height-2

	case models.ItemChangedMsg:
		m.Tick = true
		cmd = tea.Batch(m.Spinner.Tick, analyze(msg.Name))
		cmds = append(cmds, cmd)
	case models.ComponentMetaMsg:
		m.Metadata = msg.Metadata
		m.CurrentComponent = msg.Metadata.PackageName
		m.Tick = false
	case spinner.TickMsg:
		if m.Tick {
			m.Spinner, cmd = m.Spinner.Update(msg)
			cmds = append(cmds, cmd)
		}
	}
	return m, tea.Batch(cmds...)
}

func analyze(name string) tea.Cmd {
	return func() tea.Msg {
		meta, err := analyzer.AnalyzeComponent(name, ".")
		if err != nil {
			return nil
		}
		return models.ComponentMetaMsg{Metadata: meta}
	}
}
