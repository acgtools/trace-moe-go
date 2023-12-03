package tui

import (
	"github.com/acgtools/trace-moe-go/internal/search"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type searchDoneMsg bool

type errMsg struct {
	err error
}

var _ error = errMsg{}

func (e errMsg) Error() string {
	return e.err.Error()
}

type resultMsg struct {
	result *search.TraceMoeResponse
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		m.done = true
	case resultMsg:
		m.searchResult = msg
		m.done = true
	case spinner.TickMsg:
		m.spinner, cmd = m.spinner.Update(msg)
	}

	return m, cmd
}
