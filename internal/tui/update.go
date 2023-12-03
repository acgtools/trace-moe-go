package tui

import (
	"github.com/acgtools/trace-moe-go/internal/search"
	tea "github.com/charmbracelet/bubbletea"
)

type successMsg struct {
	results []*search.Result
}

type errMsg struct {
	err error
}

var _ error = errMsg{}

func (e errMsg) Error() string {
	return e.err.Error()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.size = winSize{
			width:  msg.Width,
			height: msg.Height,
		}
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case successMsg:
		m.done = true
		m.resList = newList(msg.results, m.size)
	case errMsg:
		m.err = msg.err
	}

	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	if m.done {
		m.resList, cmd = m.resList.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
