package metacomponent

import (
	"os"

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
		cmd = tea.Batch(m.Spinner.Tick, m.analyze(msg.Name))
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

func (m *Model) SetCurrentPath(path string) {
	_, err := os.Stat(path)
	if err == nil {
		m.CurrentPath = path
	}
}

func (m Model) analyze(name string) tea.Cmd {
	return func() tea.Msg {
		meta, err := analyzer.AnalyzeComponent(name, m.CurrentPath)
		if err != nil {
			return nil
		}
		return models.ComponentMetaMsg{Metadata: meta}
	}
}
